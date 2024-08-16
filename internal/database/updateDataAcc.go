package database

import (
	"database/sql"
	"fmt"
)

func UpdataDataAcc(person User, GLOBAL_PERSON User, BD_OPEN string) (bool, User) {
	fmt.Println(GLOBAL_PERSON)
	if person.Login != "" && person.Email != "" {
		db, err := sql.Open("mysql", BD_OPEN)
		if err != nil {
			panic(err)
		}

		defer db.Close()

		if person.Password != "" && person.PasswordNew != "" {
			if GLOBAL_PERSON.Password == person.Password {
				db.Query(fmt.Sprintf("UPDATE `users` SET `login` = '%s', `email` = '%s', `password` ='%s' WHERE `login` = '%s'", person.Login, person.Email, person.PasswordNew, GLOBAL_PERSON.Login))
				GLOBAL_PERSON = User{
					Login:       person.Login,
					Email:       person.Email,
					Password:    person.PasswordNew,
					PasswordNew: "",
					Success:     true,
				}
			} else {
				GLOBAL_PERSON = User{
					Login:    GLOBAL_PERSON.Login,
					Email:    GLOBAL_PERSON.Email,
					Password: GLOBAL_PERSON.Password,
					Error:    true,
					Success:  true,
				}
				return false, GLOBAL_PERSON
			}
		} else {
			db.Query(fmt.Sprintf("UPDATE `users` SET `login` = '%s', `email` = '%s' WHERE `login` = '%s'", person.Login, person.Email, GLOBAL_PERSON.Login))
			GLOBAL_PERSON = User{
				Login:    person.Login,
				Email:    person.Email,
				Password: GLOBAL_PERSON.Password,
				Success:  true,
			}
		}
		return true, GLOBAL_PERSON
	} else {
		GLOBAL_PERSON = User{
			Login:    GLOBAL_PERSON.Login,
			Email:    GLOBAL_PERSON.Email,
			Password: GLOBAL_PERSON.Password,
			Error:    true,
			Success:  true,
		}
		return false, GLOBAL_PERSON
	}
}
