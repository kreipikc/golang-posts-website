package database

import (
	"database/sql"
	"fmt"
)

// Обновление данных поста по id
func UpdatePost(BD_OPEN string, id int, post Posts) {
	db, err := sql.Open("mysql", BD_OPEN)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.Query(fmt.Sprintf("UPDATE `posts` SET `LoginAuthor` = '%s', `NamePost` = '%s', `Text` ='%s', `ImgPost` = %d WHERE `id` = %d", post.LoginAuthor, post.NamePost, post.Text, 0, id))
}
