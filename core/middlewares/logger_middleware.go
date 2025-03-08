package middlewares

import (
	"github.com/gauravsarma1992/go-rest-api/core/api"
	"go.uber.org/zap"
)

type (
	LoggerMiddleware struct {
		logger *zap.Logger
	}
)

func NewLoggerMiddleware() (loggerM *LoggerMiddleware) {
	zapLogger, _ := zap.NewProduction()
	loggerM = &LoggerMiddleware{
		logger: zapLogger,
	}
	defer zapLogger.Sync()
	return
}

func (loggerM *LoggerMiddleware) Process(req *api.Request, resp *api.Response, tr *Tracker) (err error) {
	loggerM.logger.Info("Received request",
		zap.String("RequestURI", req.RequestURI),
	)
	tr.Next()
	return
}
