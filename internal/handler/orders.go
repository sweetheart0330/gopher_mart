package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sweetheart0330/gopher_mart/internal/models"
)

func (h *Handler) DownloadOrders(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		h.log.Errorw("could not decode request body", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userLogin := r.Context().Value(models.UserIDKey).(string)
	isNew, err := h.controller.DownloadOrder(r.Context(), userLogin, order)
	if err != nil {
		if errors.Is(err, models.ErrUserConflictsId) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		if errors.Is(err, models.ErrInvalidOrderFormat) {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isNew {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	userLogin := r.Context().Value(models.UserIDKey).(string)
	orders, err := h.controller.GetOrders(r.Context(), userLogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if len(orders) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = json.NewEncoder(w).Encode(&orders)
	if err != nil {
		h.log.Errorw("could not encode request body", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
