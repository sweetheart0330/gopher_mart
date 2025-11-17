package usecase

import (
	"github.com/sweetheart0330/gopher_mart/internal/client"
	"github.com/sweetheart0330/gopher_mart/internal/domain/controller"
	"github.com/sweetheart0330/gopher_mart/internal/repository"
	"go.uber.org/zap"
)

type UseCase struct {
	log  *zap.SugaredLogger
	repo repository.IRepository
	cl   client.AccrualCalculator
}

func NewUseCase(log *zap.SugaredLogger, repo repository.IRepository, cl client.AccrualCalculator) controller.Controller {
	return &UseCase{
		log:  log,
		repo: repo,
		cl:   cl,
	}
}
