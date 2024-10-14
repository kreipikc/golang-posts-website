package database

import (
	"database/sql"
)

// Создание аккаунта
func CreatedAcc(person User, BD_OPEN string) bool {
	if person.Login != "" && person.Email != "" && person.Password != "" {
		db, err := sql.Open("mysql", BD_OPEN)
		if err != nil {
			panic(err)
		}

		defer db.Close()

		stmt, err := db.Prepare("INSERT INTO `users` (`login`, `email`, `password`) VALUES (?, ?, ?)")
		if err != nil {
			panic(err)
		}

		_, err = stmt.Exec(person.Login, person.Email, person.Password)
		if err != nil {
			panic(err)
		}
		return true
	} else {
		return false
	}
}

// Создание поста и добавление в БД
func CreatedPost(BD_OPEN string, post Posts) {
	db, _ := sql.Open("mysql", BD_OPEN)
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO `posts` (`LoginAuthor`, `NamePost`, `Text`, `ImgPost`) VALUES (?, ?, ?, ?)")
	stmt.Exec(post.LoginAuthor, post.NamePost, post.Text, post.ImgPost)
}
