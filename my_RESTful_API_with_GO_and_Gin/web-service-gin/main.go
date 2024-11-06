package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// アルバム: アメリカだと"ドル.セント"みたくなるからPriceがfloat64なようだ。
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// アルバムのスライス: 通常はDBのデータ
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

/*
エントリーポイントでルータを実行する。
注意: windowsだとcurlでなくcurl.exeだと動くようだ。
*/
func main() {
	router := gin.Default()
	/* テストコマンド
	curl.exe --url http://localhost:8080/albums `
	--header "Content-Type: application/json" `
	--include
	*/
	router.GET("/albums", getAlbums)
	/* テストコマンド
	curl.exe --request GET `
	--url http://localhost:8080/albums/2 `
	--header "Content-Type: application/json" `
	--include
	*/
	router.GET("/albums/:id", getAlbumsByID)
	/* テストコマンド
		curl.exe --request POST `
	    --url http://localhost:8080/albums `
	    --header "Content-Type: application/json" `
	    --data '{\"id\": \"4\", \"title\": \"The Modern Sound of Betty Carter\", \"artist\": \"Betty Carter\", \"price\": 49.99}' `
	    --include
	*/
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}

// アルバム全件取得
func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

// アルバム単体取得
func getAlbumsByID(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, album := range albums {
		if album.ID == id {
			ctx.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found id"})
}

// アルバム作成
func postAlbums(ctx *gin.Context) {
	var newAlbum Album
	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "body must be album's json"})
		return
	}
	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}
