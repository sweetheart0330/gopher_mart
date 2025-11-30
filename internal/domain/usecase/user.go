package usecase

import (
	"context"
	"fmt"

	"github.com/sweetheart0330/gopher_mart/internal/models"
)

func (u *UseCase) RegisterUser(ctx context.Context, user models.User) error {
	fmt.Println("RegisterUser in usecase", user)
	err := u.repo.RegisterUser(ctx, user)
	if err != nil {
		u.log.Errorw("failed to register user", "error", err)
		return err
	}

	return nil
}

func (u *UseCase) GetPasswordHash(ctx context.Context, login string) (user models.User, err error) {
	user.Password, err = u.repo.GetPasswordHash(ctx, login)
	if err != nil {
		u.log.Errorw("failed to get password", "error", err)
		return models.User{}, err
	}

	return user, nil
}
