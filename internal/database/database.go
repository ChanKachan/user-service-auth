package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

func ConnectDataBase() *pgxpool.Pool {
	connStr := "postgres://postgres:root@localhost:5432/users"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Printf("Failed to connect to database: %v\n", err)
		return nil
	}
	return dbpool
}
