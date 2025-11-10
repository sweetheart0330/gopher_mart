package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sweetheart0330/gopher_mart/internal/models"
)

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.log.Errorw("could not decode request body", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	passHash, err := h.passHasher.Hash(user.Password)
	if err != nil {
		h.log.Errorw("could not hash password", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Password = passHash

	err = h.controller.RegisterUser(user)
	if err != nil {
		if errors.Is(err, models.ErrUserAlreadyExists) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO добавить автоматическую аутентификацию

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.log.Errorw("could not decode request body", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userHash, err := h.controller.GetPasswordHash(user.Login)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	isMatch, err := h.passHasher.Verify(user.Password, userHash.Password)
	if err != nil {
		h.log.Errorw("could not verify password", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !isMatch {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	// TODO добавить cookie
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DownloadOrders(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) Withdraw(w http.ResponseWriter, r *http.Request)       {}
func (h *Handler) GetWithdrawals(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request)           {}
