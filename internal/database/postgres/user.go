package postgres

import "github.com/KokoulinM/telegram-exchange-bot/internal/models"

func (db *PostgresDatabase) CreateUser() (*models.User, error) {
	user := models.User{}

	return &user, nil
}
