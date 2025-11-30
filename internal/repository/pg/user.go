package pg

import (
	"context"
	"fmt"

	"github.com/sweetheart0330/gopher_mart/internal/models"
)

const (
	registerUserQuery   = `INSERT INTO gopher_mart.users (login, password) VALUES ($1, $2);`
	selectPasswordQuery = `SELECT password FROM gopher_mart.users WHERE users.login = $1;`
)

func (db *Database) RegisterUser(ctx context.Context, user models.User) error {
	fmt.Println("RegisterUser", user)
	_, err := db.pg.Exec(ctx, registerUserQuery, user.Login, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) GetPasswordHash(ctx context.Context, login string) (string, error) {
	err := db.pg.QueryRow(ctx, selectPasswordQuery, login).Scan(&login)
	if err != nil {
		return "", err
	}

	return login, nil
}
