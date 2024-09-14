package database

// Структура постов
type Posts struct {
	Id                          int
	LoginAuthor, NamePost, Text string
	ImgPost                     bool
}

// Структура пользователя
type User struct {
	Login, Email, Password, PasswordNew     string
	Img, Success, ErrorPassword, ErrorLogin bool
}
