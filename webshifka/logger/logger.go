package logger

import (
	"go.uber.org/zap"
)

func InitLogger() {
	logger, _ := zap.NewProduction()
}