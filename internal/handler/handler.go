package handler

import (
	"github.com/sweetheart0330/gopher_mart/internal/auth"
	"github.com/sweetheart0330/gopher_mart/internal/domain/controller"
	"go.uber.org/zap"
)

type Handler struct {
	log        zap.SugaredLogger
	passHasher auth.PassHasher
	controller controller.Controller
}

func NewHandler(l zap.SugaredLogger) *Handler {
	return &Handler{log: l}
}
