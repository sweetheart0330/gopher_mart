package controller

import (
	"context"
	"time"

	"github.com/sweetheart0330/gopher_mart/internal/models"
)

type Controller interface {
	RegisterUser(ctx context.Context, user models.User) error
	GetPasswordHash(ctx context.Context, login string) (models.User, error)

	DownloadOrder(ctx context.Context, userID string, order models.Order) (isNew bool, err error)
	GetOrders(ctx context.Context, userID string) ([]models.Order, error)

	GetBalance(ctx context.Context, userID string) (models.Balance, error)
	WithDrawn(ctx context.Context, userId string, withdraw models.Withdrawal) error
	GetWithdrawals(ctx context.Context, userID string) (withdraws []models.Withdrawal, err error)

	OrdersPooler(ctx context.Context, poolInterval time.Duration, batchSize int)
}
