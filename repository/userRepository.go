package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"regexp"
	"user-service/internal/auth"
	"user-service/internal/model"
)

type UserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserByID(id int) (*model.User, error)
	GetUserExist(user *model.User) (*model.User, error)
	//GetUserByEmail(email string) (*model.User, error)
	//UpdateUser(user *model.User) error
	//DeleteUser(id string) error
}
type userRepository struct {
	dbpool *pgxpool.Pool
	logger *zap.Logger
}

func NewUserRepository(dbpool *pgxpool.Pool, logger *zap.Logger) UserRepository {
	return &userRepository{dbpool: dbpool, logger: logger}
}

func (u *userRepository) CreateUser(user *model.User) (*model.User, error) {
	defer u.dbpool.Close()
	defer u.logger.Sync()
	var patronymic string
	var err error
	err = nil
	if &user.Patronymic == nil {
		patronymic = ""
	} else {
		patronymic = user.Patronymic
	}
	password, err := auth.HashPassword(&user.Password)
	if err != nil {
		u.logger.Error("Failed to hash password", zap.Error(err))
		return nil, err
	}
	user.Email, err = auth.IsValidEmail(user.Email)
	if err != nil {
		return nil, err
	}
	err = u.dbpool.QueryRow(context.Background(),
		`INSERT INTO "user" (name, patronymic, surname, email, phone_number, login, password) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		&user.Name, patronymic, &user.Surname, &user.Email, &user.PhoneNumber, &user.Login, &password).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, err
}
func (u *userRepository) GetUserByID(id int) (*model.User, error) {
	defer u.logger.Sync()
	defer u.dbpool.Close()
	var user model.User
	err := u.dbpool.QueryRow(context.Background(),
		`SELECT id, name, patronymic, surname, email, phone_number, login FROM "user" WHERE id = $1`, &id).
		Scan(&id, &user.Name, &user.Patronymic, &user.Surname, &user.Email, &user.PhoneNumber, &user.Login)
	if err != nil {
		return nil, err
	}
	user.Patronymic = auth.CheckPatronymic(&user.Patronymic)
	return &user, nil
}

func (u userRepository) GetUserExist(user *model.User) (*model.User, error) {
	defer u.logger.Sync()
	defer u.dbpool.Close()
	var password string
	err := u.dbpool.QueryRow(context.Background(),
		`SELECT password FROM "user" WHERE login = $1`, &user.Login).
		Scan(&password)
	if err != nil {
		return nil, err
	}

	err = auth.CheckPassword(&password, &user.Password)
	if err != nil {
		u.logger.Error("Incorrect login or password", zap.String("login", user.Login))
		return nil, err
	}
	err = u.dbpool.QueryRow(context.Background(),
		`SELECT id, name, patronymic, surname, email, phone_number, login FROM "user" WHERE login = $1`, &user.Login).
		Scan(&user.ID, &user.Name, &user.Patronymic, &user.Surname, &user.Email, &user.PhoneNumber, &user.Login)
	if err != nil {
		return nil, err
	}
	user.Patronymic = auth.CheckPatronymic(&user.Patronymic)
	return user, nil
}

func isValidEmail(email string) bool {
	// Регулярное выражение для проверки email
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Компилируем регулярное выражение
	re := regexp.MustCompile(emailRegex)

	// Проверяем, соответствует ли строка регулярному выражению
	return re.MatchString(email)
}
