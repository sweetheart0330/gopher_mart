package pg

import (
	"context"

	"github.com/sweetheart0330/gopher_mart/internal/models"
)

func (db *Database) GetBalance(ctx context.Context, userID string) (models.Balance, error) {
	//TODO implement me
	panic("implement me")
}

func (db *Database) UpdateBalance(ctx context.Context, m map[string]int) error {
	//TODO implement me
	panic("implement me")
}

func (db *Database) Withdraw(ctx context.Context, userID string, withdraw models.Withdrawal) error {
	//TODO implement me
	panic("implement me")
}

func (db *Database) GetWithdrawals(ctx context.Context, userID string) (withdraws []models.Withdrawal, err error) {
	//TODO implement me
	panic("implement me")
}
