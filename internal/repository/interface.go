package repository

import (
	"time"

	"github.com/sweetheart0330/gopher_mart/internal/models"
)

type IRepository interface {
	RegisterUser(user models.User) error
	GetPasswordHash(login string) (string, error)

	DownloadOrder(userID string, order models.Order) (bool, error)
	GetOrders(userID string) ([]models.Order, error)
}

type ISessionStore interface {
	Set(sessionID string, userID string, duration time.Duration)
	Get(string) (string, bool)
	Delete(string)
}
