package handler

import (
	"github.com/google/uuid"
	"github.com/sweetheart0330/gopher_mart/internal/auth"
	"github.com/sweetheart0330/gopher_mart/internal/domain/controller"
	"github.com/sweetheart0330/gopher_mart/internal/repository"
	"go.uber.org/zap"
)

type Handler struct {
	log          zap.SugaredLogger
	passHasher   auth.PassHasher
	controller   controller.Controller
	sessionStore repository.ISessionStore
}

func NewHandler(l zap.SugaredLogger) *Handler {
	return &Handler{log: l}
}

func (h *Handler) generateSessionID() string {
	return uuid.NewString()
}
