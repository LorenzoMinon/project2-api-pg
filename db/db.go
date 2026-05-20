package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Handler struct {
	DB *pgx.Conn
}

func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://admin:admin@localhost:5433/project2")
	return conn, err
}
