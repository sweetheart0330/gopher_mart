package handler

import "net/http"

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DownloadOrders(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) Withdraw(w http.ResponseWriter, r *http.Request)       {}
func (h *Handler) GetWithdrawals(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request)           {}
