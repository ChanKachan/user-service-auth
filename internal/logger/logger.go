package logger

import (
	"go.uber.org/zap"
	"log"
)

func InitLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Println("can't initialize zap logger: %v", err)
		return nil
	}
	return logger
}
