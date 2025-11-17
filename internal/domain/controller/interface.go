package controller

import (
	"context"
	"time"

	"github.com/sweetheart0330/gopher_mart/internal/models"
)

type Controller interface {
	RegisterUser(user models.User) error
	GetPasswordHash(login string) (models.User, error)

	DownloadOrder(userID string, order models.Order) (isNew bool, err error)
	GetOrders(userID string) ([]models.Order, error)

	GetBalance(userID string) (models.Balance, error)
	WithDrawn(userId string, withdraw models.Withdrawal) error

	OrdersPooler(ctx context.Context, poolInterval time.Duration, batchSize int)
}
