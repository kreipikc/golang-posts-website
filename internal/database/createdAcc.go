package database

import (
	"database/sql"
)

type User struct {
	Login, Email, Password, PasswordNew string
	Success, Error                      bool
}

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
