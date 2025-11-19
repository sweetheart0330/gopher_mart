package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/sweetheart0330/gopher_mart/internal/auth"
	httpAccr "github.com/sweetheart0330/gopher_mart/internal/client/http"
	"github.com/sweetheart0330/gopher_mart/internal/config"
	"github.com/sweetheart0330/gopher_mart/internal/domain/usecase"
	"github.com/sweetheart0330/gopher_mart/internal/handler"
	"github.com/sweetheart0330/gopher_mart/internal/repository/pg"
	"github.com/sweetheart0330/gopher_mart/internal/repository/session"
	"github.com/sweetheart0330/gopher_mart/internal/router"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func RunServer(ctx context.Context) error {
	opt, err := config.GetOptions()
	if err != nil {
		return fmt.Errorf("failed to get options: %w", err)
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		return fmt.Errorf("failed to init logger, err: %w", err)
	}

	defer logger.Sync()
	sugar := *logger.Sugar()

	repo, err := pg.NewDatabase(ctx, opt.URIDatabase, &sugar)
	if err != nil {
		return fmt.Errorf("failed to init database, err: %w", err)
	}

	cl := httpAccr.NewClient(opt.AccrualAddr)
	u := usecase.NewUseCase(&sugar, repo, cl)

	passHasher := auth.NewPasswordHasher(opt.AuthPepper, opt.AuthCost)
	sessStore := session.NewStore()
	h := handler.NewHandler(&sugar, passHasher, u, sessStore)
	route := router.NewRouter(*h)

	eg, egCtx := errgroup.WithContext(ctx)

	server := &http.Server{
		Addr:    opt.Host,
		Handler: route,
	}

	eg.Go(func() error {
		sugar.Infow("Starting server", "cfg", opt.Host)
		if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("failed to start server: %w", err)
		}

		return nil
	})

	eg.Go(func() error {
		<-egCtx.Done()
		shCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		sugar.Infow("Stopping server", "cfg", opt.Host)

		return server.Shutdown(shCtx)
	})

	return eg.Wait()
}
