package worker

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/muammarahlnn/learnyscape-backend/gateway/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/gateway/internal/log"
	"github.com/muammarahlnn/learnyscape-backend/pkg/logger"
	"github.com/spf13/cobra"
)

func Start() {
	cfg := config.InitConfig()
	log.SetLogger(logger.NewZeroLogLogger(cfg.Logger.Level))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	rootCmd := &cobra.Command{}
	cmd := []*cobra.Command{
		{
			Use:   "serve-all",
			Short: "Run all",
			Run: func(cmd *cobra.Command, args []string) {
				runHttpWorker(ctx, cfg)
			},
		},
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Logger.Fatal(err)
	}
}
