package models

import "time"

const (
	UserIDKey = "user-id" // Ключ для контекста
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Order struct {
	OrderID    string    `json:"order_id"`
	UserID     int       `json:"user_id"`
	Status     string    `json:"status"`
	Accrual    int       `json:"accrual"`
	UploadedAt time.Time `json:"uploaded_at"`
}

type Balance struct {
	UserID     int `json:"user_id"`
	BalanceSum int `json:"balance"`
	WithDrawn  int `json:"with_drawn"`
}

type Withdrawal struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Order       string    `json:"order"`
	Sum         int       `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}
