package log

import "github.com/muammarahlnn/learnyscape-backend/pkg/logger"

var Logger logger.Logger

func SetLogger(logger logger.Logger) {
	Logger = logger
}
