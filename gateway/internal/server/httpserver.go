package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/learnyscape-backend/gateway/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/gateway/internal/log"
	"github.com/muammarahlnn/learnyscape-backend/gateway/internal/provider"
	"github.com/muammarahlnn/learnyscape-backend/pkg/middleware"
)

type HttpServer struct {
	cfg *config.Config
	srv *http.Server
}

func NewHttpServer(cfg *config.Config) *HttpServer {
	gin.SetMode(cfg.App.Envinronment)

	router := gin.New()
	router.ContextWithFallback = true
	router.HandleMethodNotAllowed = true

	registerMiddleware(router, cfg)
	provider.BootstrapHttp(cfg, router)

	return &HttpServer{
		cfg: cfg,
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.HttpServer.Host, cfg.HttpServer.Port),
			Handler: router,
		},
	}
}

func (s *HttpServer) Start() {
	log.Logger.Info("Running HTTP server on port:", s.cfg.HttpServer.Port)

	err := s.srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Logger.Fatal("Error while HTTP server listening:", err)
	}

	log.Logger.Info("HTTP server is not receiving new requests...")
}

func (s *HttpServer) Shutdown() {
	timeout := time.Duration(s.cfg.HttpServer.GracePeriod) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Logger.Info("Attempting to shutdown HTTP server...")
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Logger.Fatal("Error shutting down HTTP server:", err)
	}

	log.Logger.Info("HTTP server shutdown gracefully")
}

func registerMiddleware(router *gin.Engine, cfg *config.Config) {
	middlewares := []gin.HandlerFunc{
		gin.Recovery(),
		middleware.LoggerMiddleware(log.Logger),
		gzip.Gzip(gzip.BestSpeed),
		middleware.ErrorMiddleware(),
		middleware.TimeoutCancelMiddleware(cfg.HttpServer.RequestTimeoutPeriod),
		cors.New(cors.Config{
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			AllowAllOrigins:  true,
			AllowCredentials: true,
		}),
	}

	router.Use(middlewares...)
}
