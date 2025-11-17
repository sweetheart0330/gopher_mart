package models

import "errors"

var (
	ErrUserAlreadyExists = errors.New("User already exists")
	ErrUserNotFound      = errors.New("User not found")
	// ErrUserConflictsId - не описывает четко проблему, чтобы злоумышленники не могли воспользоваться описанием ошибки
	ErrUserConflictsId    = errors.New("ID cannot be added, watch logs for more information")
	ErrInvalidOrderFormat = errors.New("Invalid order number format")

	ErrTooManyRequests   = errors.New("Too many requests")
	ErrInsufficientFunds = errors.New("Insufficient funds")
)
