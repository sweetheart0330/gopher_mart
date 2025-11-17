package client

import "github.com/sweetheart0330/gopher_mart/internal/models"

type AccrualCalculator interface {
	GetCalcOrder(orderNumber string) (*models.Order, error)
}
