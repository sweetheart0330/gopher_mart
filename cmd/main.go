package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/sweetheart0330/gopher_mart/internal/app"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			stop()
			return nil
		case <-egCtx.Done():
			return nil
		}
	})

	eg.Go(func() error { return app.RunServer(egCtx) })

	if err := eg.Wait(); err != nil {
		log.Println("Error running server", zap.Error(err))
		return
	}

	log.Println("Server exited")
}
