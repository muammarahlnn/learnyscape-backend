package worker

import (
	"context"

	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/server"
)

func runHttpWorker(ctx context.Context, cfg *config.Config) {
	srv := server.NewHttpServer(cfg)
	go srv.Start()

	<-ctx.Done()
	srv.Shutdown()
}
