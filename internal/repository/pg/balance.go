package pg

import "github.com/sweetheart0330/gopher_mart/internal/models"

func (db *Database) GetBalance(userID string) (models.Balance, error) {
	//TODO implement me
	panic("implement me")
}

func (db *Database) UpdateBalance(m map[string]int) error {
	//TODO implement me
	panic("implement me")
}

func (db *Database) Withdraw(userID string, withdraw models.Withdrawal) error {
	//TODO implement me
	panic("implement me")
}

func (db *Database) GetWithdrawals(userID string) (withdraws []models.Withdrawal, err error) {
	//TODO implement me
	panic("implement me")
}
