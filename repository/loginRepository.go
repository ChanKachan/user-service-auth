package repository

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"time"
	"user-service/internal/auth"
	"user-service/internal/config"
	"user-service/internal/model"
)

type LoginRepository interface {
	Authenticate(Login *model.LoginRequest)
}

type loginRepository struct {
	dbpool *pgxpool.Pool
	logger *zap.Logger
}

func NewLoginRepository(dbpool *pgxpool.Pool, logger *zap.Logger) *loginRepository {
	return &loginRepository{dbpool: dbpool, logger: logger}
}

func (repo *loginRepository) Authenticate(login *model.LoginRequest) (model.LoginResponse, error) {
	defer repo.dbpool.Close()
	defer repo.logger.Sync()

	var id int
	var name, surname, patronymic, phone, email, loginUser, password, role string
	var authInfo model.LoginResponse

	role = ""
	err := repo.dbpool.QueryRow(context.Background(),
		`SELECT id,name, surname, patronymic,phone_number,email,login, password FROM "user" WHERE login = $1`, &login.Login).
		Scan(&id, &name, &surname, &patronymic, &phone, &email, &loginUser, &password, &loginUser, &password)
	if err != nil {
		err = repo.dbpool.QueryRow(context.Background(),
			`SELECT e.id,e.name, e.surname, e.patronymic,e.phone_number,e.email,e.login, e.password, jt.job_title FROM "employee" e INNER JOIN "job_title" jt ON e.id_job_title = jt.id  WHERE e.login = $1`, &login.Login).
			Scan(&id, &name, &surname, &patronymic, &phone, &email, &loginUser, &password, &role)
		if err != nil {
			err = errors.New("login or password is incorrect")
			return authInfo, err
		}
	}

	err = auth.CheckPassword(&password, &login.Password)
	if err != nil {
		repo.logger.Error("Incorrect login or password", zap.String("login", login.Login))
		return authInfo, err
	}

	payload := jwt.MapClaims{
		"sub":  login.Login,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
		"iat":  time.Now().Unix(),
		"role": role,
	}

	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(config.JWTSecretKey)
	if err != nil {
		err = errors.New("failed to create token JWT")
		return authInfo, err
	}

	user := model.Employee{
		ID:          id,
		JobTitle:    "",
		Email:       email,
		Name:        name,
		Surname:     surname,
		Patronymic:  patronymic,
		PhoneNumber: phone,
		Login:       loginUser,
		Password:    "",
	}
	if role != "" {
		user.JobTitle = role
	}
	authInfo = model.LoginResponse{AccessToken: t, User: user}
	return authInfo, nil

}
