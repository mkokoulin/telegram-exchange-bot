package postgres

import "github.com/mkokoulin/telegram-exchange-bot/internal/models"

func (db *PostgresDatabase) CreateUser() (*models.User, error) {
	user := models.User{}

	return &user, nil
}
