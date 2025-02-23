package model

type User struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	PhoneNumber string `json:"phone_number"`
	Login       string `json:"login"`
	Password    string `json:"password"`
}
