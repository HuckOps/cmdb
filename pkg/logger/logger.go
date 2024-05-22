package logger

import (
	"go.uber.org/zap"
)

var ServerLogger *zap.Logger

func InitLogger() {
	ServerLogger, _ = zap.NewProduction()

}
