package transport

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	database "mymodule.com/v2/internal/database"
)

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
	person := database.User{
		Login:    r.FormValue("login"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
		Success:  true,
	}

	checkForVoid := database.CreatedAcc(person, BD_OPEN)
	if checkForVoid {
		t := template.Must(template.ParseFiles("../../web/templates/index.html"))
		t.Execute(w, person)
	} else {
		fmt.Println("Форма имеет пустые значения.")
	}

}

func enter_to_acc(w http.ResponseWriter, r *http.Request) {
	person := database.User{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
		Success:  false,
	}

	person, existence := database.CheckUserInBD(person, BD_OPEN)
	if existence {
		t := template.Must(template.ParseFiles("../../web/templates/index.html"))
		t.Execute(w, person)
	} else {
		t := template.Must(template.ParseFiles("../../web/templates/authorization.html"))
		t.Execute(w, true)
	}
}

func Handlefunc() {
	InitConfig()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../web/static"))))
	http.HandleFunc("/index", index)
	http.HandleFunc("/registration", registration)
	http.HandleFunc("/authorization", authorization)
	http.HandleFunc("/enter_to_acc", enter_to_acc)
	http.HandleFunc("/created_acc", created_acc)
	http.ListenAndServe(PORT, nil)
}
