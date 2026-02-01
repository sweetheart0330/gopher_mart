package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sweetheart0330/gopher_mart/internal/models"
)

const (
	SelectOrders = `
		SELECT o.id, o.order_id, o.user_id, o.status, o.accrual, o.uploaded_at
		FROM gopher_mart.orders AS o
		WHERE user_id = (SELECT users.id FROM gopher_mart.users WHERE login = $1)
`
	InsertOrder = `
    	INSERT INTO gopher_mart.orders (order_id, user_id, status, accrual, uploaded_at)
    	VALUES ($1, (SELECT id FROM gopher_mart.users WHERE login = $2), $3, $4, $5)
    	ON CONFLICT (order_id) DO NOTHING
    	RETURNING id
`
	SelectNotCalculatedOrders = `
		SELECT o.id, o.order_id, o.user_id, o.status, o.accrual, o.uploaded_at 
		FROM gopher_mart.orders AS o
		WHERE (o.status = 'NEW' OR o.status = 'PROCESSING');
`
	UpdateOrder = `
		UPDATE gopher_mart.orders AS o
        SET 
            status = $1,
            accrual = $2
        WHERE o.order_id = $3`
)

func (db *Database) DownloadOrder(ctx context.Context, userLogin string, order models.Order) (bool, error) {
	_, err := db.pg.Exec(ctx, InsertOrder, order.OrderID, userLogin, order.Status, order.Accrual, order.UploadedAt)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (db *Database) GetOrders(ctx context.Context, userLogin string) ([]models.Order, error) {
	rows, err := db.pg.Query(ctx, SelectOrders, userLogin)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err = rows.Scan(
			&order.ID,
			&order.OrderID,
			&order.UserID,
			&order.Status,
			&order.Accrual,
			&order.UploadedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (db *Database) GetNotCalcOrders(ctx context.Context) ([]models.Order, error) {
	rows, err := db.pg.Query(ctx, SelectNotCalculatedOrders)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err = rows.Scan(
			&order.ID,
			&order.OrderID,
			&order.UserID,
			&order.Status,
			&order.Accrual,
			&order.UploadedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (db *Database) updateOrders(ctx context.Context, tx pgx.Tx, orders []models.Order) error {
	for _, order := range orders {
		_, err := tx.Exec(ctx, UpdateOrder, order.Status, order.Accrual, order.OrderID)
		if err != nil {
			return fmt.Errorf("failed to send request: %w", err)
		}
	}

	return nil
}
