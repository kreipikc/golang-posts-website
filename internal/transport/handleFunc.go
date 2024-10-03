package transport

import (
	"net/http"

	database "mymodule.com/v2/internal/database"
	servies "mymodule.com/v2/internal/servies"
)

var GLOBAL_PERSON database.User         // Глобальная структура Пользователя
var MAP_LIST_IMG = servies.InitMapImg() // Глобальная мапа, которая хранит все картинки

// Обработка всех страниц
func Handlefunc() {
	InitConfig()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
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
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/about_us", about_us)
	http.HandleFunc("/post_page", post_page)
	http.HandleFunc("/created_post", created_post)
	http.HandleFunc("/deleted_post", deleted_post)
	http.HandleFunc("/settings_post", settings_post)
	http.HandleFunc("/update_post", update_post)
	http.ListenAndServe(":"+PORT, nil)
}
