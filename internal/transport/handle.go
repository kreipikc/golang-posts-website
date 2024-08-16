package transport

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	database "mymodule.com/v2/internal/database"
)

var GLOBAL_PERSON database.User

// Основная страница
func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../../web/templates/index.html"))
	t.Execute(w, GLOBAL_PERSON)
}

// Страница с формой регистрацией
func registration(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../../web/templates/registration.html"))
	t.Execute(w, nil)
}

// Страница с формой авторизацией
func authorization(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../../web/templates/authorization.html"))
	t.Execute(w, nil)
}

// Обработка создания нового пользователя
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

// Обработка запроса на вход в аккаунт
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

// Страница настроек аккаунта пользователя
func settings_user(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../../web/templates/settings_user.html"))
	t.Execute(w, GLOBAL_PERSON)
}

// Выход из аккаунта
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

// Обновление данных из settings_user
func update_user(w http.ResponseWriter, r *http.Request) {
	person := database.User{
		Login:       r.FormValue("login"),
		Email:       r.FormValue("email"),
		Password:    r.FormValue("password_old"),
		PasswordNew: r.FormValue("password_new"),
	}
	check, person_new := database.UpdataDataAcc(person, GLOBAL_PERSON, BD_OPEN)
	GLOBAL_PERSON = person_new
	if check {
		t := template.Must(template.ParseFiles("../../web/templates/settings_user.html"))
		t.Execute(w, GLOBAL_PERSON)
	} else {
		t := template.Must(template.ParseFiles("../../web/templates/settings_user.html"))
		t.Execute(w, GLOBAL_PERSON)
	}
}

// Обработка
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
