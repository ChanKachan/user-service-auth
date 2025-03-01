package auth

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"reflect"
	"user-service/internal/model"
)

func HashPassword(password *string) (*string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	str := string(bytes)
	return &str, err
}

// CheckPatronymic проверяет является ли patronymic nil или нет
func CheckPatronymic(patronymic *string) string {
	if patronymic == nil {
		*patronymic = ""
		return *patronymic
	} else {
		return *patronymic
	}
}

// CheckPassword проверяет, соответствует ли введенный пароль хешу.
func CheckPassword(hashedPassword, password *string) error {
	return bcrypt.CompareHashAndPassword([]byte(*hashedPassword), []byte(*password))
}

// IsValidEmail проверяет действительность адреса электронной почты.
func IsValidEmail(email string) (string, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return "", err
	}
	return email, nil
}

func GetAllJsonTeg(user *model.User) []string {
	v := reflect.ValueOf(user)
	t := v.Type()
	tagsArr := make([]string, 0)
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json") // Получаем значение тега "json"
		tagsArr = append(tagsArr, tag)
		fmt.Printf("%s: %s\n", field.Name, tag)
	}
	return tagsArr
}
