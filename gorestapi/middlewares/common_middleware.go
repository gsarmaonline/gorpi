package middlewares

import (
	"container/list"

	"github.com/gauravsarma1992/go-rest-api/gorestapi/api"
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

func (loggerM *LoggerMiddleware) Process(req *api.Request, lElem *list.Element) (resp *api.Response, err error) {
	loggerM.logger.Info("Received request",
		zap.String("RequestURI", req.RequestURI),
	)
	lElem.Next()
	return
}
