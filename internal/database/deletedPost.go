package database

import (
	"database/sql"
	"fmt"
)

func DeletedPost(BD_OPEN string, id int) {
	db, _ := sql.Open("mysql", BD_OPEN)
	defer db.Close()

	db.Query(fmt.Sprintf("DELETE FROM `posts` WHERE `ID` = %d", id))
}
