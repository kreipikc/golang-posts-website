package database

import "database/sql"

// Создание поста и добавление в БД
func CreatedPost(BD_OPEN string, post Posts) {
	db, _ := sql.Open("mysql", BD_OPEN)
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO `posts` (`LoginAuthor`, `NamePost`, `Text`, `ImgPost`) VALUES (?, ?, ?, ?)")
	stmt.Exec(post.LoginAuthor, post.NamePost, post.Text, post.ImgPost)
}
