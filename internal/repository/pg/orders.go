package pg

import "github.com/sweetheart0330/gopher_mart/internal/models"

func (db *Database) DownloadOrder(userID string, order models.Order) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (db *Database) GetOrders(userID string) ([]models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (db *Database) GetNotCalcOrders() ([]models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (db *Database) UpdateOrders(orders []models.Order) error {
	//TODO implement me
	panic("implement me")
}
