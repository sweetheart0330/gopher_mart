package models

import "time"

const (
	UserIDKey = "user-id" // Ключ для контекста
)

type OrderStatus string

const (
	New        OrderStatus = "NEW"
	Processing OrderStatus = "PROCESSING"
	Invalid    OrderStatus = "INVALID"
	Processed  OrderStatus = "PROCESSED"
)

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Order struct {
	ID         int         `json:"id"`
	OrderID    string      `json:"order_id"`
	UserID     int         `json:"user_id"`
	Status     OrderStatus `json:"status"`
	Accrual    int         `json:"accrual"`
	UploadedAt time.Time   `json:"uploaded_at"`
}

type Balance struct {
	ID         int `json:"id"`
	UserID     int `json:"user_id"`
	BalanceSum int `json:"balance"`
	WithDrawn  int `json:"with_drawn"`
}

type Withdrawal struct {
	ID          int       `json:"-"`
	WithdrawID  int       `json:"order"`
	UserId      int       `json:"-"`
	Sum         int       `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}
