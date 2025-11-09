package handler

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

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
