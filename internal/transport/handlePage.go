package transport

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	database "mymodule.com/v2/internal/database"
)

// Основная страница
func index(w http.ResponseWriter, r *http.Request) {
	GLOBAL_PERSON.ErrorLogin = false
	GLOBAL_PERSON.ErrorPassword = false
	t := template.Must(template.ParseFiles("../../web/templates/index.html"))
	t.Execute(w, GLOBAL_PERSON)
}

// Страница с формой регистрации
func registration(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../../web/templates/registration.html"))
	t.Execute(w, nil)
}

// Страница с формой авторизации
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

	checkForLogin, _ := database.CheckUserInBDLogin(person, BD_OPEN)
	checkForVoid := database.CreatedAcc(person, BD_OPEN)

	// Если пользователь ввел все данные
	if checkForVoid {
		// Если такого логина ЕЩЁ нет в БД (нет такого логина в БД)
		if !checkForLogin {
			t := template.Must(template.ParseFiles("../../web/templates/index.html"))
			GLOBAL_PERSON = person
			t.Execute(w, GLOBAL_PERSON)
		} else {
			person.ErrorLogin = true
			GLOBAL_PERSON = person
			t := template.Must(template.ParseFiles("../../web/templates/registration.html"))
			t.Execute(w, GLOBAL_PERSON)
		}
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

	// Если данные верны
	if existence {
		t := template.Must(template.ParseFiles("../../web/templates/index.html"))
		GLOBAL_PERSON = person

		// Если у пользователя есть загруженная картинка на сервере
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
	GLOBAL_PERSON.ErrorPassword = false
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

	http.Redirect(w, r, "/authorization", http.StatusSeeOther)
}

// Обновление данных из settings_user
func update_user(w http.ResponseWriter, r *http.Request) {
	person := database.User{
		Login:       r.FormValue("login"),
		Email:       r.FormValue("email"),
		Password:    r.FormValue("password_old"),
		PasswordNew: r.FormValue("password_new"),
	}

	checkForLogin, tempLogin := database.CheckUserInBDLogin(person, BD_OPEN)

	// Если такого логина ЕЩЁ нет в БД (нет такого логина в БД); Исключение: Используется старый логин, то есть не меняется в настройках при сохранении
	if !checkForLogin || (checkForLogin && tempLogin == GLOBAL_PERSON.Login) {
		check, person_new := database.UpdataDataAcc(person, GLOBAL_PERSON, BD_OPEN)

		// Проверка на наличие аватарки по старому имени
		if _, ok := MAP_LIST_IMG[fmt.Sprintf("%s.jpg", GLOBAL_PERSON.Login)]; ok {
			os.Rename(fmt.Sprintf("../../web/static/img/profile_img/%s.jpg", GLOBAL_PERSON.Login), fmt.Sprintf("../../web/static/img/profile_img/%s.jpg", person_new.Login))
			MAP_LIST_IMG[fmt.Sprintf("%s.jpg", person_new.Login)] = true // Добовляем новую картинку в мапу
			delete(MAP_LIST_IMG, GLOBAL_PERSON.Login)                    // Удаляем старую картинку из мапы
			person_new.Img = true
		}

		GLOBAL_PERSON = person_new

		// Если всё прошло успешно (в database.UpdataDataAcc)
		if check {
			t := template.Must(template.ParseFiles("../../web/templates/settings_user.html"))
			t.Execute(w, GLOBAL_PERSON)
		} else {
			t := template.Must(template.ParseFiles("../../web/templates/settings_user.html"))
			t.Execute(w, GLOBAL_PERSON)
		}
	} else {
		GLOBAL_PERSON.ErrorLogin = true
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

	// Если файл был передан (не был пустым при сохранении)
	if fileHeader != nil {
		defer file.Close()
		contentType := fileHeader.Header["Content-Type"][0]
		var osFile *os.File

		// Если файл был передан в формате .jpg или .png
		if contentType == "image/jpeg" || contentType == "image/png" {
			osFile, _ = os.Create(fmt.Sprintf("../../web/static/img/profile_img/%s.jpg", GLOBAL_PERSON.Login))
			defer osFile.Close()

			fileBytes, _ := io.ReadAll(file)
			osFile.Write(fileBytes)

			MAP_LIST_IMG[fmt.Sprintf("%s.jpg", GLOBAL_PERSON.Login)] = true
			GLOBAL_PERSON.Img = true
		}

		t := template.Must(template.ParseFiles("../../web/templates/settings_img.html"))
		t.Execute(w, GLOBAL_PERSON)
	} else {
		t := template.Must(template.ParseFiles("../../web/templates/settings_img.html"))
		t.Execute(w, GLOBAL_PERSON)
	}
}
