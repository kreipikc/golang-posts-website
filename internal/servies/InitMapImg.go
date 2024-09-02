package servies

import (
	"os"
)

// Инициализируем мапу, обратывая папку с аватарками
func InitMapImg() map[string]bool {
	var map_list = map[string]bool{"": false}
	files, _ := os.ReadDir("web/static/img/profile_img")
	for _, file := range files {
		map_list[file.Name()] = true
	}
	return map_list
}
