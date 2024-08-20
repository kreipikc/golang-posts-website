package database

import (
	"database/sql"
	"fmt"
)

func CheckUserInBDLogin(person User, BD_OPEN string) bool {
	if person.Login != "" {
		db, _ := sql.Open("mysql", BD_OPEN)
		defer db.Close()

		res, _ := db.Query(fmt.Sprintf("SELECT * FROM `users` WHERE `login` = '%s'", person.Login))

		for res.Next() {
			var us User
			_ = res.Scan(&us.Login, &us.Email, &us.Password)
			if us.Login == person.Login {
				return true
			}
		}
	}
	return false
}
