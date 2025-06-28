package worker

import (
	"context"

	"github.com/muammarahlnn/learnyscape-backend/verification-service/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/verification-service/internal/server"
)

func runKafkaWorker(ctx context.Context, cfg *config.Config) {
	srv := server.NewKafkaServer(cfg)
	go srv.Start()

	<-ctx.Done()
	srv.Shutdown()
}
