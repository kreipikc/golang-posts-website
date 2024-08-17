package database

import (
	"database/sql"
	"fmt"
)

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
