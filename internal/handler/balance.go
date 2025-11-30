package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sweetheart0330/gopher_mart/internal/models"
)

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(models.UserIDKey).(string)

	balance, err := h.controller.GetBalance(r.Context(), userId)
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

func (h *Handler) Withdraw(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(models.UserIDKey).(string)

	var withdraw models.Withdrawal
	err := json.NewDecoder(r.Body).Decode(&withdraw)
	if err != nil {
		h.log.Errorw("could not decode request body", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.controller.WithDrawn(r.Context(), userId, withdraw)
	if err != nil {
		if errors.Is(err, models.ErrInsufficientFunds) {
			http.Error(w, err.Error(), http.StatusPaymentRequired)
			return
		}

		if errors.Is(err, models.ErrUserConflictsId) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetWithdrawals(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(models.UserIDKey).(string)

	resp, err := h.controller.GetWithdrawals(r.Context(), userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(resp) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		h.log.Errorw("could not encode response", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
