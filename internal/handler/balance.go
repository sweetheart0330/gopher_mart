package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sweetheart0330/gopher_mart/internal/models"
)

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(models.UserIDKey).(string)

	balance, err := h.controller.GetBalance(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&balance)
	if err != nil {
		h.log.Errorw("could not encode request body", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
