package servies

import (
	"fmt"
	"os"

	"mymodule.com/v2/internal/database"
)

// Проверка на наличие картинки в локальной директории
func CheckOnDir(person database.User) bool {
	files, _ := os.ReadDir("../../web/static/img/profile_img")
	for _, file := range files {
		if fmt.Sprintf("%s.jpg", person.Login) == file.Name() {
			return true
		}
	}
	return false
}
