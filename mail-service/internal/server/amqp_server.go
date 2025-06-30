package server

import (
	"context"
	"time"

	"github.com/muammarahlnn/learnyscape-backend/mail-service/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/mail-service/internal/log"
	"github.com/muammarahlnn/learnyscape-backend/mail-service/internal/provider"
	"github.com/muammarahlnn/learnyscape-backend/pkg/mq"
)

type AMQPServer struct {
	ctx       context.Context
	cancel    context.CancelFunc
	consumers []mq.AMQPConsumer
}

func NewAMQPServer(cfg *config.Config) *AMQPServer {
	ctx, cancel := context.WithCancel(context.Background())
	return &AMQPServer{
		ctx:       ctx,
		cancel:    cancel,
		consumers: provider.BootstrapAMQP(cfg),
	}
}

func (s *AMQPServer) Start() {
	const numberOfWorkers = 3

	log.Logger.Info("Running AMQP server...")
	for _, consumer := range s.consumers {
		go consumer.Consume(s.ctx, numberOfWorkers)
	}
	log.Logger.Info("AMQP server is running...")
}

func (s *AMQPServer) Shutdown() {
	log.Logger.Info("Attempting to shutdown AMQP server...")
	time.Sleep(10 * time.Second)

	s.cancel()
	for _, consumer := range s.consumers {
		if err := consumer.Close(); err != nil {
			log.Logger.Errorf("Error closing consumer %s: %s", consumer.Queue(), err)
		}
	}

	log.Logger.Info("AMQP server shut down gracefully")
}
