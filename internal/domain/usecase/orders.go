package usecase

import (
	"strconv"

	"github.com/sweetheart0330/gopher_mart/internal/models"
)

func (u *UseCase) DownloadOrder(userID string, order models.Order) (bool, error) {
	if !Valid(order.OrderID) {
		u.log.Errorw("invalid order id", "order_id", order.OrderID)
		return false, models.ErrInvalidOrderFormat
	}

	isNew, err := u.repo.DownloadOrder(userID, order)
	if err != nil {
		u.log.Errorw("failed to download order",
			"order_id", order.OrderID,
			"err", err)
		return false, err
	}

	return isNew, nil
}

func (u *UseCase) GetOrder(userID string) ([]models.Order, error) {
	orders, err := u.repo.GetOrders(userID)
	if err != nil {
		u.log.Errorw("could not get orders",
			"userID", userID,
			"err", err)
		return nil, err
	}

	return orders, nil
}

func Valid(number string) bool {
	num, err := strconv.Atoi(number)
	if err != nil {

		return false
	}
	return (num%10+checksum(num/10))%10 == 0
}

func checksum(number int) int {
	var luhn int

	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 { // even
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		luhn += cur
		number = number / 10
	}
	return luhn % 10
}
