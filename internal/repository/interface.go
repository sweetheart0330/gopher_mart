package repository

import (
	"context"
	"time"

	"github.com/sweetheart0330/gopher_mart/internal/models"
)

type IRepository interface {
	RegisterUser(ctx context.Context, user models.User) error
	GetPasswordHash(ctx context.Context, login string) (string, error)

	DownloadOrder(ctx context.Context, userID string, order models.Order) (bool, error)
	GetOrders(ctx context.Context, userID string) ([]models.Order, error)
	GetNotCalcOrders(ctx context.Context) ([]models.Order, error)

	GetBalance(ctx context.Context, userID string) (models.Balance, error)
	UpdatePollerStatuses(ctx context.Context, orders []models.Order, balance map[int]int) error

	Withdraw(ctx context.Context, userLogin string, withdraw models.Withdrawal) error
	GetWithdrawals(ctx context.Context, userID string) (withdraws []models.Withdrawal, err error)
}

type ISessionStore interface {
	Set(sessionID string, userID string, duration time.Duration)
	Get(string) (string, bool)
	Delete(string)
}
