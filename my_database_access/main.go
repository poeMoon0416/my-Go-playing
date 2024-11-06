package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// モデル
type Album struct {
	ID     int
	Title  string
	Artist string
	Price  float64
}

// どの関数からでもアクセスできるようにグローバル変数にしているが実際はあまりよくないようだ
var db *sql.DB

func main() {
	// DSNの定義
	cfg := mysql.Config{
		// クレデンシャルはハードコーディングを防ぐため環境変数にする(実行前にset)
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "my_Go_playing",
	}

	// DB接続
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("fail to connect MySQL server: %v", err)
	}

	// DB接続の死活監視(接続できなくなった場合再接続を試みるようだ)
	if err := db.Ping(); err != nil {
		log.Fatalf("fail to ping MySQL server: %v", err)
	}

	if albums, err := albumsByArtist("John Coltrane"); err != nil {
		log.Fatalf("fail to exec albumsByArtist: %v", err)
	} else {
		fmt.Println(albums)
	}

	if album, err := albumByID(3); err != nil {
		log.Fatalf("fail to exec albumByID: %v", err)
	} else {
		fmt.Println(album)
	}

	if id, err := createAlbum(Album{Title: "マツケンサンバ", Artist: "マツケン", Price: 12.34}); err != nil {
		log.Fatalf("fail to exec createAlbum: %v", err)
	} else {
		fmt.Println(id)
	}
}

// アーティストでアルバム検索(複数行)
func albumsByArtist(artist string) ([]Album, error) {
	// 複数行が返るクエリを実行
	rows, err := db.Query("SELECT * FROM albums WHERE artist = ?", artist)
	if err != nil {
		return nil, fmt.Errorf("fail to exec query: %v", err)
	}
	// 複数行では一時ファイルに結果行が保存されるのかもしれない(一時ファイルのClose()と考えると辻褄は合う)
	defer rows.Close()

	// Goの可変長引数は1つ以上の引数を表すようだ
	albums := make([]Album, 0)
	/*
		rows.Next()は行を読み取る準備を行う(該当の1行をオンメモリにしているのかも)
		最初の読取でも実行する必要がある
		行の読取に成功した場合はtrue, 行を読み切ったかエラーが起きた場合はfalse
	*/
	for rows.Next() {
		var album Album
		// 行の読取と構造体への割り当て
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, fmt.Errorf("fail to bind columns to struct: %v", err)
		}
		albums = append(albums, album)
	}

	// rows.Next()がエラーで抜けてきた場合
	if rows.Err() != nil {
		return nil, fmt.Errorf("fail to read rows")
	}

	return albums, nil
}

// IDでアルバム検索(単一行)
func albumByID(id int) (Album, error) {
	var album Album
	// 単一行を返すクエリはエラーをrow.Scan()のタイミングで処理するようだ
	row := db.QueryRow("SELECT * FROM albums WHERE id = ?", id)
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		return Album{}, fmt.Errorf("not exists id: %v", err)
	}
	return album, nil
}

// アルバム作成(結果行なし、実行結果の詳細情報)
func createAlbum(album Album) (int, error) {
	// 結果行を返さないクエリは実行結果の詳細情報が返る
	res, err := db.Exec("INSERT INTO albums(title, artist, price) VALUES(?, ?, ?)", album.Title, album.Artist, album.Price)
	if err != nil {
		return 0, fmt.Errorf("fail to create album: %v", err)
	}

	// int64のAUTO_INCREMENTのIDとerrorが返る
	id, insErr := res.LastInsertId()
	if insErr != nil {
		return 0, fmt.Errorf("fail to get last insert id: %v", err)
	}
	return int(id), insErr
}
