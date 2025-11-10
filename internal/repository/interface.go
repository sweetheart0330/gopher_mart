package repository

import "github.com/sweetheart0330/gopher_mart/internal/models"

type IRepository interface {
	RegisterUser(user models.User) error
	GetPasswordHash(login string) (string, error)
}
