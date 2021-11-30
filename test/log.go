package main

import (
	"go.uber.org/zap"
	"time"
)


func main(){
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "11111"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	zap.New()
}
