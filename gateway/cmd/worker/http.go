package worker

import (
	"context"

	"github.com/muammarahlnn/learnyscape-backend/gateway/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/gateway/internal/server"
)

func runHttpWorker(ctx context.Context, cfg *config.Config) {
	srv := server.NewHttpServer(cfg)
	go srv.Start()

	<-ctx.Done()
	srv.Shutdown()
}
