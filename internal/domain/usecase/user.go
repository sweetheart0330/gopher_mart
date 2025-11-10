package usecase

import (
	"github.com/sweetheart0330/gopher_mart/internal/models"
)

func (u *UseCase) RegisterUser(user models.User) error {
	err := u.repo.RegisterUser(user)
	if err != nil {
		u.log.Errorw("failed to register user", "error", err)
		return err
	}

	return nil
}

func (u *UseCase) GetPasswordHash(login string) (user models.User, err error) {
	user.Password, err = u.repo.GetPasswordHash(login)
	if err != nil {
		u.log.Errorw("failed to get password", "error", err)
		return models.User{}, err
	}

	return user, nil
}
