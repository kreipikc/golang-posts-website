package database

import (
	"database/sql"
	"fmt"
)

// Инициализирование всех постов из БД
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

// Проверка пользователя на наличие в БД
func CheckUserInBD(person User, BD_OPEN string) (User, bool) {
	if person.Login != "" && person.Password != "" {
		db, _ := sql.Open("mysql", BD_OPEN)
		defer db.Close()

		res, _ := db.Query(fmt.Sprintf("SELECT * FROM `users` WHERE `login` = '%s'", person.Login))

		for res.Next() {
			var us User
			_ = res.Scan(&us.Login, &us.Email, &us.Password)
			if us.Login == person.Login && us.Password == person.Password {
				person = us
				break
			}
		}

		if person.Email != "" {
			person.Success = true
			return person, true
		}
	}
	return person, false
}

// Проверка на существование пользователя с таким же логином
func CheckUserInBDLogin(person User, BD_OPEN string) (bool, string) {
	if person.Login != "" {
		db, _ := sql.Open("mysql", BD_OPEN)
		defer db.Close()

		res, _ := db.Query(fmt.Sprintf("SELECT * FROM `users` WHERE `login` = '%s'", person.Login))

		for res.Next() {
			var us User
			_ = res.Scan(&us.Login, &us.Email, &us.Password)
			if us.Login == person.Login {
				return true, us.Login
			}
		}
	}
	return false, ""
}
