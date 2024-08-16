package transport

import "os"

// Глобальные перменные конфига
var (
	PORT    string
	BD_OPEN string
)

// Инициализация данных конфига
func InitConfig() {
	P, err := os.ReadFile("../../internal/config/port.txt")
	if err != nil {
		panic(err)
	}
	PORT = string(P)

	B, err := os.ReadFile("../../internal/config/bdOpen.txt")
	if err != nil {
		panic(err)
	}
	BD_OPEN = string(B)
}
