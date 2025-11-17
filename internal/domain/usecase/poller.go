package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sweetheart0330/gopher_mart/internal/models"
)

func (u *UseCase) OrdersPooler(ctx context.Context, poolInterval time.Duration, batchSize int) {
	ticker := time.NewTicker(poolInterval)
	defer ticker.Stop()
	for {
	L:
		select {
		case <-ctx.Done():
			u.log.Info("Stopping OrdersPooler")
			return
		case <-ticker.C:
			orders, err := u.repo.GetNotCalcOrders()
			if err != nil {
				u.log.Errorw("failed to get orders", "err", err)
			}

			// TODO добавить retry-механизм
			// TODO проверить корректность работы цикла
			for i := 0; i < len(orders); i += batchSize {
				err = u.updateOrderWithAccrual(orders[i : i+batchSize])
				if err != nil {
					u.log.Errorw("failed to update orders batch", "err", err)
					continue L
				}

				if i+batchSize >= len(orders) {
					batchSize = len(orders) - 1 - i
				}
			}
		}
	}
}

func (u *UseCase) updateOrderWithAccrual(orders []models.Order) error {
	userTotalSum := make(map[string]int)

	for i := 0; i < len(orders); {
		err := u.getUpdatedOrder(&orders[i])
		if err != nil {
			if errors.Is(err, models.ErrTooManyRequests) {
				u.log.Warnw("too many requests", "error", err, "order", orders[i].OrderID)

				continue
			}

			return fmt.Errorf("failed to update order status with accrual: %w", err)
		}

		if orders[i].Status == "SUCCESS" {
			userTotalSum[orders[i].OrderID] += orders[i].Accrual
		}

		i++
	}

	err := u.repo.UpdateOrders(orders)
	if err != nil {
		return fmt.Errorf("failed to update orders in repo: %w", err)
	}

	err = u.repo.UpdateBalance(userTotalSum)
	if err != nil {
		return fmt.Errorf("failed to update balance in repo: %w", err)
	}

	return nil
}

func (u *UseCase) getUpdatedOrder(order *models.Order) error {
	resp, err := u.cl.GetCalcOrder(order.OrderID)
	if err != nil {
		return err
	}

	order.Accrual = resp.Accrual
	order.UploadedAt = resp.UploadedAt
	order.Status = resp.Status

	return nil
}
