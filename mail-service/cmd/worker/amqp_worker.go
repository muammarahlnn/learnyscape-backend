package worker

import (
	"context"

	"github.com/muammarahlnn/learnyscape-backend/mail-service/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/mail-service/internal/server"
)

func runAMQPWorker(cfg *config.Config, ctx context.Context) {
	srv := server.NewAMQPServer(cfg)
	go srv.Start()

	<-ctx.Done()
	srv.Shutdown()
}
