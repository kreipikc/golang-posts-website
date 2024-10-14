package transport

import (
	"net/http"

	"github.com/gorilla/mux"
	database "mymodule.com/v2/internal/database"
	servies "mymodule.com/v2/internal/servies"
)

var GLOBAL_PERSON database.User         // Глобальная структура Пользователя
var MAP_LIST_IMG = servies.InitMapImg() // Глобальная мапа, которая хранит все картинки

// Обработка всех страниц
func Handlefunc() {
	InitConfig()
	router := mux.NewRouter()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	router.HandleFunc("/update_post/{id:[0-9+]}", update_post)
	router.HandleFunc("/index", index)
	router.HandleFunc("/registration", registration)
	router.HandleFunc("/authorization", authorization)
	router.HandleFunc("/enter_to_acc", enter_to_acc)
	router.HandleFunc("/created_acc", created_acc)
	router.HandleFunc("/settings_user", settings_user)
	router.HandleFunc("/exit_acc", exit_acc)
	router.HandleFunc("/update_user", update_user)
	router.HandleFunc("/settings_img", settings_img)
	router.HandleFunc("/update_img", update_img)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/about_us", about_us)
	router.HandleFunc("/post_page", post_page)
	router.HandleFunc("/created_post", created_post)
	router.HandleFunc("/deleted_post", deleted_post)
	router.HandleFunc("/settings_post", settings_post)
	http.Handle("/", router)
	http.ListenAndServe(":"+PORT, nil)
}
