package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func JobTitleToId(db *pgxpool.Pool, jobTitle string) (*int, error) {
	var id int
	contx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := db.QueryRow(contx, `select id from "job_title" where job_title = $1`, jobTitle).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
