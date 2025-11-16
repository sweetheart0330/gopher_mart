package models

import "errors"

var (
	ErrUserAlreadyExists = errors.New("User already exists")
	ErrUserNotFound      = errors.New("User not found")
	// ErrUserConflictsOrderId - не описывает четко проблему, чтобы злоумышленники не могли воспользоваться описанием ошибки
	ErrUserConflictsOrderId = errors.New("Order ID cannot be added, watch logs for more information")
	ErrInvalidOrderFormat   = errors.New("Invalid order number format")
)
