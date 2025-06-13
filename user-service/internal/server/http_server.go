package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/muammarahlnn/learnyscape-backend/pkg/middleware"
	validationutil "github.com/muammarahlnn/learnyscape-backend/pkg/util/validation"
	"github.com/muammarahlnn/user-service/internal/config"
	"github.com/muammarahlnn/user-service/internal/log"
	"github.com/muammarahlnn/user-service/internal/provider"
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

	registerValidators()
	registerMiddleware(cfg, router)

	provider.BootstrapHttp(router)

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

func registerValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(validationutil.TagNameFormatter)
		v.RegisterValidation("password", validationutil.Password)
	}
}

func registerMiddleware(cfg *config.Config, router *gin.Engine) {
	middlewares := []gin.HandlerFunc{
		gin.Recovery(),
		middleware.LoggerMiddleware(log.Logger),
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
