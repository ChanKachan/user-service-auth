package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"time"
	"user-service/internal/auth"
	"user-service/internal/model"
)

type EmployeeRepository interface {
	CreateEmployee(employee model.Employee) (string, error)
}

type employeeRepository struct {
	dbpool *pgxpool.Pool
	logger *zap.Logger
}

func NewEmployeeRepository(dbpool *pgxpool.Pool, logger *zap.Logger) EmployeeRepository {
	return &employeeRepository{dbpool: dbpool, logger: logger}
}

func (e employeeRepository) CreateEmployee(employee model.Employee) (string, error) {
	defer e.dbpool.Close()
	defer e.logger.Sync()
	var err error

	contx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	role := *auth.Role

	if role != "administrator" {
		err = errors.New("У вас нет прав администратора, чтобы выполнить данное действие")
		return "", err
	}

	id, err := JobTitleToId(e.dbpool, employee.JobTitle)
	if err != nil {
		return "", err
	}

	patronymic := auth.CheckPatronymic(&employee.Patronymic)
	password, err := auth.HashPassword(&employee.Password)
	if err != nil {
		return "", err
	}

	_, err = e.dbpool.Exec(contx,
		`INSERT INTO "employee" (id_job_title, name, patronymic, surname, email, phone_number, login, password) VALUES($1, $2, $3, $4, $5, $6, $7, $8)`,
		&id, &employee.Name, patronymic, &employee.Surname, &employee.Email, &employee.PhoneNumber, &employee.Login, &password)
	if err != nil {
		return "", err
	}
	return "The employee has been successfully added to the database.", nil
}

// ======
