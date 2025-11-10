package controller

import "github.com/sweetheart0330/gopher_mart/internal/models"

type Controller interface {
	RegisterUser(user models.User) error
	GetPasswordHash(login string) (models.User, error)
}
