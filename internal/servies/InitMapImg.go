package servies

import (
	"os"
)

// Инициализируем папку с аватарками и записываем в глобальную переменную
func InitMapImg() map[string]bool {
	var map_list = map[string]bool{"": false}
	files, _ := os.ReadDir("../../web/static/img/profile_img")
	for _, file := range files {
		map_list[file.Name()] = true
	}
	return map_list
}
