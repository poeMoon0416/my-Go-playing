DROP TABLE IF EXISTS albums;

-- MySQLでは慣例的にPRIMARY KEYにはNOT NULLを書くことも多いようだ。
CREATE TABLE albums (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    price DECIMAL(5, 2) NOT NULL
);

INSERT INTO albums(title, artist, price) VALUES
('Blue Train', 'John Coltrane', 56.99),
('Giant Steps', 'John Coltrane', 63.99),
('Jeru', 'Gerry Mulligan', 17.99),
('Sarah Vaughan', 'Sarah Vaughan', 34.98);
-- source setup_albums.sql