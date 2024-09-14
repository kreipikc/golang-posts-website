package database

import (
	"database/sql"
	"fmt"
)

// Инициализация поста по id
func CheckPostsSolo(BD_OPEN string, id int) Posts {
	db, _ := sql.Open("mysql", BD_OPEN)
	defer db.Close()

	res, _ := db.Query(fmt.Sprintf("SELECT * FROM `posts` WHERE `id` = %d", id))
	var post Posts
	for res.Next() {
		_ = res.Scan(&post.Id, &post.LoginAuthor, &post.NamePost, &post.Text, &post.ImgPost)
	}
	return post
}
