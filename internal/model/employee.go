package model

type Employee struct {
	ID          int    `json:"id"`
	JobTitle    string `json:"job_title"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	PhoneNumber string `json:"phone_number"`
	Login       string `json:"login"`
	Password    string `json:"password"`
}
