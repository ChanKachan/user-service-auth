package auth

import (
	"golang.org/x/crypto/bcrypt"
	"net/mail"
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
