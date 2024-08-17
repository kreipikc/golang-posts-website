package transport

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	database "mymodule.com/v2/internal/database"
	servies "mymodule.com/v2/internal/servies"
)

var GLOBAL_PERSON database.User
var MAP_LIST_IMG = map[string]bool{"": false}

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

		// Инициализируем папку с аватарками и записываем в глобальную переменную
		temp := servies.InitMapImg()
		MAP_LIST_IMG = temp

		if MAP_LIST_IMG[fmt.Sprintf("%s.jpg", GLOBAL_PERSON.Login)] {
			GLOBAL_PERSON.Img = true
		}
		t.Execute(w, GLOBAL_PERSON)
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
	// Проверка на наличие аватарки по старому имени
	if _, ok := MAP_LIST_IMG[fmt.Sprintf("%s.jpg", GLOBAL_PERSON.Login)]; ok {
		os.Rename(fmt.Sprintf("../../web/static/img/profile_img/%s.jpg", GLOBAL_PERSON.Login), fmt.Sprintf("../../web/static/img/profile_img/%s.jpg", person_new.Login))
		MAP_LIST_IMG[fmt.Sprintf("%s.jpg", person_new.Login)] = true
		person_new.Img = true
	}
	GLOBAL_PERSON = person_new
	if check {
		t := template.Must(template.ParseFiles("../../web/templates/settings_user.html"))
		t.Execute(w, GLOBAL_PERSON)
	} else {
		t := template.Must(template.ParseFiles("../../web/templates/settings_user.html"))
		t.Execute(w, GLOBAL_PERSON)
	}
}

// Страница для настройки аватарки
func settings_img(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../../web/templates/settings_img.html"))
	t.Execute(w, GLOBAL_PERSON)
}

// Обновление картинки из settings_img
func update_img(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10)
	file, fileHeader, _ := r.FormFile("file_input")
	if fileHeader != nil {
		defer file.Close()
		contentType := fileHeader.Header["Content-Type"][0]
		var osFile *os.File
		if contentType == "image/jpeg" || contentType == "image/png" {
			osFile, _ = os.Create(fmt.Sprintf("../../web/static/img/profile_img/%s.jpg", GLOBAL_PERSON.Login))
		}
		defer osFile.Close()
		fileBytes, _ := io.ReadAll(file)
		osFile.Write(fileBytes)
		MAP_LIST_IMG[fmt.Sprintf("%s.jpg", GLOBAL_PERSON.Login)] = true

		GLOBAL_PERSON.Img = true
	}
	t := template.Must(template.ParseFiles("../../web/templates/settings_img.html"))
	t.Execute(w, GLOBAL_PERSON)
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
	http.HandleFunc("/settings_img", settings_img)
	http.HandleFunc("/update_img", update_img)
	http.ListenAndServe(PORT, nil)
}
