package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sweetheart0330/gopher_mart/internal/handler"
)

func NewRouter(h handler.Handler) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(h.MiddlewareLogger())
	//mux.Use(h.PoolWorker())
	//mux.Use(h.DecompressHandle)
	//mux.Use(h.CompressHandle)
	//mux.Use(h.CheckHashSum)

	mux.Route("/api/user", func(r chi.Router) {
		r.Post("/register", h.RegisterUser)
		r.Post("/login", h.LoginUser)
		r.Route("/orders", func(ro chi.Router) {
			ro.Post("/", h.DownloadOrders)
			ro.Get("/", h.GetOrders)
		})
		r.Route("/balance", func(rb chi.Router) {
			rb.Get("/", h.GetBalance)
			rb.Post("/withdraw", h.Withdraw)
		})
		r.Get("/withdrawals", h.GetWithdrawals)
	})
	mux.Get("/ping", h.Ping)

	return mux
}
