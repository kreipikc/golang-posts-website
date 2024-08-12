package transport

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	Login, Email, Password string
	Success                bool
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../../web/templates/index.html"))
	t.Execute(w, nil)
}

func registration(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../../web/templates/registration.html"))
	t.Execute(w, nil)
}

func authorization(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../../web/templates/authorization.html"))
	t.Execute(w, nil)
}

func created_acc(w http.ResponseWriter, r *http.Request) {
	person := user{
		Login:    r.FormValue("login"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
		Success:  true,
	}

	if person.Login != "" && person.Email != "" && person.Password != "" {
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go-reg")
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

		t := template.Must(template.ParseFiles("../../web/templates/index.html"))
		t.Execute(w, person)
	} else {
		fmt.Println("Форма имеет пустые значения.")
	}
}

func enter_to_acc(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	if login != "" && password != "" {
		db, _ := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go-reg")
		defer db.Close()

		res, _ := db.Query("SELECT * FROM `users`")

		test := false
		for res.Next() {
			var us user
			_ = res.Scan(&us.Login, &us.Email, &us.Password)
			if us.Login == login && us.Password == password {
				test = true
			}
		}

		if test {
			person := user{
				Login:    r.FormValue("login"),
				Email:    r.FormValue("email"),
				Password: r.FormValue("password"),
				Success:  true,
			}
			t := template.Must(template.ParseFiles("../../web/templates/index.html"))
			t.Execute(w, person)
		} else {
			t := template.Must(template.ParseFiles("../../web/templates/authorization.html"))
			t.Execute(w, true)
		}
	}
}

func Handlefunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../web/static"))))
	http.HandleFunc("/index", index)
	http.HandleFunc("/registration", registration)
	http.HandleFunc("/authorization", authorization)
	http.HandleFunc("/enter_to_acc", enter_to_acc)
	http.HandleFunc("/created_acc", created_acc)
	http.ListenAndServe(":8080", nil)
}
