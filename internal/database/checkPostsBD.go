package database

import "database/sql"

type Posts struct {
	Id                          int
	LoginAuthor, NamePost, Text string
	ImgPost                     bool
}

func CheckPosts(BD_OPEN string) []Posts {
	db, _ := sql.Open("mysql", BD_OPEN)
	defer db.Close()

	res, _ := db.Query("SELECT * FROM `posts`")

	var posts []Posts
	for res.Next() {
		var post Posts
		_ = res.Scan(&post.Id, &post.LoginAuthor, &post.NamePost, &post.Text, &post.ImgPost)
		posts = append(posts, post)
	}
	return posts
}
