package pg

import (
	"context"

	"github.com/sweetheart0330/gopher_mart/internal/models"
)

const (
	SelectOrders = `
		SELECT o.id, o.order_id, o.user_id, o.status, o.accrual, o.uploaded 
		FROM gopher_mart.orders AS o
		WHERE user_id = (SELECT users.id FROM gopher_mart.users WHERE login = $1)
`
	InsertOrder = `
    	INSERT INTO gopher_mart.orders (order_id, user_id, status, accrual, uploaded)
    	VALUES ($1, (SELECT id FROM gopher_mart.users WHERE login = $2), $3, $4, $5)
    	ON CONFLICT (order_id) DO NOTHING
    	RETURNING id
`
	SelectNotCalculatedOrders = `
		SELECT o.id, o.order_id, o.user_id, o.status, o.accrual, o.uploaded 
		FROM gopher_mart.orders AS o
		WHERE o.status = 'NEW' || o.status = 'PROCESSING';
`
	UpdateOrders = `
		UPDATE gopher_mart.orders AS o
        SET 
            status = u.status,
            accrual = u.accrual
        FROM (
            SELECT 
                unnest($1::text[]) AS order_id,
                unnest($2::text[]) AS status,
                unnest($3::numeric[]) AS accrual
        ) AS u
        WHERE o.order_id = u.order_id`
)

func (db *Database) DownloadOrder(ctx context.Context, userID string, order models.Order) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (db *Database) GetOrders(ctx context.Context, userID string) ([]models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (db *Database) GetNotCalcOrders(ctx context.Context) ([]models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (db *Database) UpdateOrders(ctx context.Context, orders []models.Order) error {
	//TODO implement me
	panic("implement me")
}
