package postgres

import (
	"database/sql"

	"github.com/rs/zerolog"
)

type PostgresDatabase struct {
	conn   *sql.DB
	logger *zerolog.Logger
}

func New(conn *sql.DB, logger *zerolog.Logger) *PostgresDatabase {
	return &PostgresDatabase{
		conn:   conn,
		logger: logger,
	}
}
