package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/sweetheart0330/gopher_mart/internal/models"
)

func (h *Handler) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userID, valid := h.sessionStore.Get(cookie.Value)
		if !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), models.UserIDKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Handler) MiddlewareLogger() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			reqID := middleware.GetReqID(r.Context())
			defer func() {
				h.log.Infow(
					"REQUEST COMPLETED",
					"reqID", reqID,
					"method", r.Method,
					"path", r.URL.Path,
					"status", ww.Status(),
					"duration", time.Since(t1).String(),
				)
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}
