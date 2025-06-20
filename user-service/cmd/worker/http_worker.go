package worker

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/muammarahlnn/learnyscape-backend/pkg/logger"
	"github.com/muammarahlnn/user-service/internal/config"
	"github.com/muammarahlnn/user-service/internal/log"
	"github.com/muammarahlnn/user-service/internal/provider"
	"github.com/spf13/cobra"
)

func Start() {
	cfg := config.InitConfig()
	log.SetLogger(logger.NewZapLogger(cfg.Logger.Level))
	provider.BootstrapGlobal(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	rootCmd := &cobra.Command{}
	cmds := []*cobra.Command{
		{
			Use:   "serve-all",
			Short: "Run all",
			Run: func(cmd *cobra.Command, args []string) {
				runHttpWorker(ctx, cfg)
			},
		},
	}

	rootCmd.AddCommand(cmds...)
	if err := rootCmd.Execute(); err != nil {
		log.Logger.Fatal(err)
	}
}
