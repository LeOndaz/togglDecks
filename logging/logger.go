package logging

import (
	"go.uber.org/zap"
	"log"
)

func initLogger() *zap.Logger {
	logger, err := zap.NewProduction()

	if err != nil {
		// a better scenario is falling back to go builtin logger
		log.Fatal("couldn't initialize logger, existing")
	}

	return logger
}

var Logger = initLogger()
