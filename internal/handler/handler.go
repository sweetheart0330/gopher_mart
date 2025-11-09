package handler

import "go.uber.org/zap"

type Handler struct {
	log zap.SugaredLogger
}

func NewHandler(l zap.SugaredLogger) *Handler {
	return &Handler{log: l}
}
