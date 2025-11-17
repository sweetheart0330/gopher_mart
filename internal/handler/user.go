package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

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

	h.setCookie(w, user.Login)

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

	h.setCookie(w, user.Login)

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) setCookie(w http.ResponseWriter, userLogin string) {
	sessID := h.generateSessionID()

	h.sessionStore.Set(sessID, userLogin, 24*time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_cookie",
		Value:    sessID,
		Path:     "/",
		MaxAge:   86400,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
}
