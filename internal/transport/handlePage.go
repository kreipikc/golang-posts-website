package transport

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	database "mymodule.com/v2/internal/database"
)

// Структура для передаваемых данных
type IndexData struct {
	Person           database.User    // Данные о пользователе
	Posts            []database.Posts // Массив постов
	Post             database.Posts   // Пост в единичном экземпляре
	CreatedOrSetting bool             // Тумблер. true - created post; false - setting post
}

// Основная страница
func index(w http.ResponseWriter, r *http.Request) {
	GLOBAL_PERSON.ErrorLogin = false
	GLOBAL_PERSON.ErrorPassword = false
	posts := database.CheckPosts(BD_OPEN)
	data := IndexData{
		Person: GLOBAL_PERSON,
		Posts:  posts,
	}
	t := template.Must(template.ParseFiles("web/templates/index.html"))
	t.Execute(w, data)
}

// Обработка страницы контактов
func contact(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("web/templates/contact.html"))
	t.Execute(w, GLOBAL_PERSON)
}

// Обработка страницы about_us
func about_us(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("web/templates/about_us.html"))
	t.Execute(w, GLOBAL_PERSON)
}

// Страница с формой регистрации
func registration(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("web/templates/registration.html"))
	t.Execute(w, nil)
}

// Страница с формой авторизации
func authorization(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("web/templates/authorization.html"))
	t.Execute(w, nil)
}

// Обработка создания нового пользователя
func created_acc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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
				posts := database.CheckPosts(BD_OPEN)
				GLOBAL_PERSON = person
				data := IndexData{
					Person: GLOBAL_PERSON,
					Posts:  posts,
				}
				t := template.Must(template.ParseFiles("web/templates/index.html"))
				t.Execute(w, data)
			} else {
				person.ErrorLogin = true
				GLOBAL_PERSON = person
				t := template.Must(template.ParseFiles("web/templates/registration.html"))
				t.Execute(w, GLOBAL_PERSON)
			}
		} else {
			fmt.Println("Форма имеет пустые значения.")
		}
	}
}

// Обработка запроса на вход в аккаунт
func enter_to_acc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		person := database.User{
			Login:    r.FormValue("login"),
			Password: r.FormValue("password"),
			Success:  false,
		}

		person, existence := database.CheckUserInBD(person, BD_OPEN)

		// Если данные верны
		if existence {
			GLOBAL_PERSON = person
			posts := database.CheckPosts(BD_OPEN)

			// Если у пользователя есть загруженная картинка на сервере
			if MAP_LIST_IMG[fmt.Sprintf("%s.jpg", GLOBAL_PERSON.Login)] {
				GLOBAL_PERSON.Img = true
			}

			data := IndexData{
				Person: GLOBAL_PERSON,
				Posts:  posts,
			}
			t := template.Must(template.ParseFiles("web/templates/index.html"))
			t.Execute(w, data)
		} else {
			t := template.Must(template.ParseFiles("web/templates/authorization.html"))
			t.Execute(w, true)
		}
	}
}

// Страница настроек аккаунта пользователя
func settings_user(w http.ResponseWriter, r *http.Request) {
	GLOBAL_PERSON.ErrorPassword = false
	t := template.Must(template.ParseFiles("web/templates/settings_user.html"))
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
	if r.Method == "POST" {
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
				os.Rename(fmt.Sprintf("web/static/img/profile_img/%s.jpg", GLOBAL_PERSON.Login), fmt.Sprintf("web/static/img/profile_img/%s.jpg", person_new.Login))
				MAP_LIST_IMG[fmt.Sprintf("%s.jpg", person_new.Login)] = true     // Добовляем новую картинку в мапу
				delete(MAP_LIST_IMG, fmt.Sprintf("%s.jpg", GLOBAL_PERSON.Login)) // Удаляем старую картинку из мапы
				person_new.Img = true
			}

			GLOBAL_PERSON = person_new

			// Если всё прошло успешно (в database.UpdataDataAcc)
			if check {
				t := template.Must(template.ParseFiles("web/templates/settings_user.html"))
				t.Execute(w, GLOBAL_PERSON)
			} else {
				t := template.Must(template.ParseFiles("web/templates/settings_user.html"))
				t.Execute(w, GLOBAL_PERSON)
			}
		} else {
			GLOBAL_PERSON.ErrorLogin = true
			t := template.Must(template.ParseFiles("web/templates/settings_user.html"))
			t.Execute(w, GLOBAL_PERSON)
		}
	}
}

// Страница для настройки аватарки
func settings_img(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("web/templates/settings_img.html"))
	t.Execute(w, GLOBAL_PERSON)
}

// Обновление картинки из settings_img
func update_img(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(10)
		file, fileHeader, _ := r.FormFile("file_input")

		// Если файл был передан (не был пустым при сохранении)
		if fileHeader != nil {
			defer file.Close()
			contentType := fileHeader.Header["Content-Type"][0]
			var osFile *os.File

			// Если файл был передан в формате .jpg или .png
			if contentType == "image/jpeg" || contentType == "image/png" {
				osFile, _ = os.Create(fmt.Sprintf("web/static/img/profile_img/%s.jpg", GLOBAL_PERSON.Login))
				defer osFile.Close()

				fileBytes, _ := io.ReadAll(file)
				osFile.Write(fileBytes)

				MAP_LIST_IMG[fmt.Sprintf("%s.jpg", GLOBAL_PERSON.Login)] = true
				GLOBAL_PERSON.Img = true
			}

			t := template.Must(template.ParseFiles("web/templates/settings_img.html"))
			t.Execute(w, GLOBAL_PERSON)
		} else {
			t := template.Must(template.ParseFiles("web/templates/settings_img.html"))
			t.Execute(w, GLOBAL_PERSON)
		}
	}
}

// Страница для создания постов
func post_page(w http.ResponseWriter, r *http.Request) {
	data := IndexData{
		Person:           GLOBAL_PERSON,
		CreatedOrSetting: true,
	}
	t := template.Must(template.ParseFiles("web/templates/post_page.html"))
	t.Execute(w, data)
}

// Запрос на создание поста
func created_post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		post := database.Posts{
			LoginAuthor: GLOBAL_PERSON.Login,
			NamePost:    r.FormValue("namePost"),
			Text:        r.FormValue("textPost"),
			ImgPost:     false,
		}

		database.CreatedPost(BD_OPEN, post)
		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}

// Удаление поста
func deleted_post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("number2")
		idInt, _ := strconv.Atoi(id)
		database.DeletedPost(BD_OPEN, idInt)
		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}

// Страница для обновления данных (настройки) поста
func settings_post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("number1")
		id_int, _ := strconv.Atoi(id)

		post := database.CheckPostsSolo(BD_OPEN, id_int)

		data := IndexData{
			Person:           GLOBAL_PERSON,
			CreatedOrSetting: false,
			Post:             post,
		}
		t := template.Must(template.ParseFiles("web/templates/post_page.html"))
		t.Execute(w, data)
	}
}

// Запрос на обновление данных поста
func update_post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("number")
		id_int, _ := strconv.Atoi(id)
		post := database.Posts{
			LoginAuthor: GLOBAL_PERSON.Login,
			NamePost:    r.FormValue("namePost"),
			Text:        r.FormValue("textPost"),
			ImgPost:     false,
		}

		database.UpdatePost(BD_OPEN, id_int, post)
		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}
