package worker

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/muammarahlnn/learnyscape-backend/mail-service/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/mail-service/internal/log"
	"github.com/muammarahlnn/learnyscape-backend/mail-service/internal/provider"
	"github.com/muammarahlnn/learnyscape-backend/pkg/logger"
	"github.com/spf13/cobra"
)

func Start() {
	cfg := config.InitConfig()
	log.SetLogger(logger.NewZeroLogLogger(cfg.Logger.Level))
	provider.BootstrapGlobal(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	rootCmd := &cobra.Command{}
	cmds := []*cobra.Command{
		{
			Use:   "serve-all",
			Short: "Run all",
			Run: func(cmd *cobra.Command, _ []string) {
				runAMQPWorker(cfg, ctx)
			},
		},
	}

	rootCmd.AddCommand(cmds...)
	if err := rootCmd.Execute(); err != nil {
		log.Logger.Fatal(err)
	}
}
