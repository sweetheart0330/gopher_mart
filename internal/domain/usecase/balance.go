package usecase

import (
	"context"

	"github.com/sweetheart0330/gopher_mart/internal/models"
)

func (u *UseCase) GetBalance(ctx context.Context, userId string) (balance models.Balance, err error) {
	balance, err = u.repo.GetBalance(userId)
	if err != nil {
		u.log.Errorw("failed to get balance from repo", "err", err)

		return models.Balance{}, err
	}

	return balance, nil
}

func (u *UseCase) WithDrawn(ctx context.Context, userId string, withdraw models.Withdrawal) error {
	err := u.repo.Withdraw(userId, withdraw)
	if err != nil {
		u.log.Errorw("failed to withdraw from repo", "err", err)
		return err
	}

	return nil
}

func (u *UseCase) GetWithdrawals(ctx context.Context, userID string) (withdraws []models.Withdrawal, err error) {
	withdrawals, err := u.GetWithdrawals(ctx, userID)
	if err != nil {
		u.log.Errorw("failed to get withdrawals from repo", "err", err)
		return nil, err
	}

	return withdrawals, nil
}
