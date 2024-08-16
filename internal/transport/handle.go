package transport

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	database "mymodule.com/v2/internal/database"
)

var GLOBAL_PERSON database.User

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../../web/templates/index.html"))
	t.Execute(w, GLOBAL_PERSON)
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
		GLOBAL_PERSON = person
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
		GLOBAL_PERSON = person
		t.Execute(w, person)
	} else {
		t := template.Must(template.ParseFiles("../../web/templates/authorization.html"))
		t.Execute(w, true)
	}
}

func settings_user(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../../web/templates/settings_user.html"))
	t.Execute(w, GLOBAL_PERSON)
}

func exit_acc(w http.ResponseWriter, r *http.Request) {
	GLOBAL_PERSON = database.User{
		Login:    "",
		Email:    "",
		Password: "",
		Success:  false,
	}
	t := template.Must(template.ParseFiles("../../web/templates/authorization.html"))
	t.Execute(w, nil)
}

func update_user(w http.ResponseWriter, r *http.Request) {
	person := database.User{
		Login:       r.FormValue("login"),
		Email:       r.FormValue("email"),
		Password:    r.FormValue("password_old"),
		PasswordNew: r.FormValue("password_new"),
	}
	fmt.Println(person)
	check, GLOBAL_PERSON := database.UpdataDataAcc(person, GLOBAL_PERSON, BD_OPEN)
	fmt.Println("gl", GLOBAL_PERSON)
	if check {
		t := template.Must(template.ParseFiles("../../web/templates/settings_user.html"))
		t.Execute(w, GLOBAL_PERSON)
	} else {
		t := template.Must(template.ParseFiles("../../web/templates/settings_user.html"))
		t.Execute(w, GLOBAL_PERSON)
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
	http.HandleFunc("/settings_user", settings_user)
	http.HandleFunc("/exit_acc", exit_acc)
	http.HandleFunc("/update_user", update_user)
	http.ListenAndServe(PORT, nil)
}
