package handlers

import (
	"github.com/jackc/pgx/v5"
)

type Handler struct {
	DB *pgx.Conn
}
