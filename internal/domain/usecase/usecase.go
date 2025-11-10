package usecase

import (
	"github.com/sweetheart0330/gopher_mart/internal/repository"
	"go.uber.org/zap"
)

type UseCase struct {
	log  *zap.SugaredLogger
	repo repository.IRepository
}
