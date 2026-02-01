package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sweetheart0330/gopher_mart/internal/models"
)

const (
	selectBalance = `
	SELECT b.id, b.user_id, b.balance, b.with_drawn
	FROM gopher_mart.balances AS b
	WHERE user_id = (SELECT users.id FROM gopher_mart.users WHERE login = $1);
`
	updateBalance = `
	INSERT INTO gopher_mart.balances (user_id, balance)
	VALUES ($1, $2)
	ON CONFLICT (user_id)
	DO UPDATE SET
    	balance = gopher_mart.balances.balance + EXCLUDED.balance;
`
	insertWithDrawn = `
	INSERT INTO gopher_mart.withdrawals (withdraw_id, user_id, sum)
	VALUES ($1, $2, $3);
`
	selectUserWithDrawals = `
	SELECT w.id, w.withdraw_id, w.user_id, w.sum, w.processed_at 
	FROM gopher_mart.withdrawals AS w
	WHERE user_id = (SELECT users.id FROM gopher_mart.users WHERE login = $1);

`
)

func (db *Database) GetBalance(ctx context.Context, userID string) (models.Balance, error) {
	var b models.Balance
	err := db.pg.QueryRow(ctx, selectBalance, userID).Scan(&b.ID, &b.UserID, &b.BalanceSum, &b.WithDrawn)
	if err != nil {
		return models.Balance{}, err
	}

	return b, nil
}

func (db *Database) UpdatePollerStatuses(ctx context.Context, orders []models.Order, balance map[int]int) error {
	tx, err := db.pg.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	err = db.updateBalance(ctx, tx, balance)
	if err != nil {
		return fmt.Errorf("failed to update balance: %v", err)
	}

	err = db.updateOrders(ctx, tx, orders)
	if err != nil {
		return fmt.Errorf("failed to update orders: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (db *Database) updateBalance(ctx context.Context, tx pgx.Tx, balance map[int]int) error {
	for userId, accrual := range balance {
		_, err := tx.Exec(ctx, updateBalance, userId, accrual)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *Database) Withdraw(ctx context.Context, userLogin string, withdraw models.Withdrawal) error {
	tx, err := db.pg.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	var b models.Balance
	err = tx.QueryRow(ctx, selectBalance, userLogin).Scan(&b.ID, &b.UserID, &b.BalanceSum, &b.WithDrawn)
	if err != nil {
		return fmt.Errorf("failed to get balance: %w", err)
	}

	withdraw.UserId = b.UserID
	if b.BalanceSum < withdraw.Sum {
		return models.ErrInsufficientFunds
	}

	_, err = tx.Exec(ctx, insertWithDrawn, withdraw.WithdrawID, withdraw.UserId, withdraw.Sum)
	if err != nil {
		return fmt.Errorf("failed to insert withdraw: %w", err)
	}

	_, err = tx.Exec(ctx, updateBalance, b.UserID, withdraw.Sum*(-1))
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil
}

func (db *Database) GetWithdrawals(ctx context.Context, userLogin string) (withdraws []models.Withdrawal, err error) {

	rows, err := db.pg.Query(ctx, selectUserWithDrawals, userLogin)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	for rows.Next() {
		var wd models.Withdrawal
		err = rows.Scan(
			&wd.ID,
			&wd.WithdrawID,
			&wd.UserId,
			&wd.Sum,
			&wd.ProcessedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan withdrawal: %w", err)
		}

		withdraws = append(withdraws, wd)
	}

	return withdraws, nil
}
