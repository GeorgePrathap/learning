package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	logger.Debug("User logged in", zap.String("username", "john.doe"))
	logger.Info("Application started", zap.String("version", "1.0.0"))
	logger.Warn("Disk space running low", zap.Float64("usage", 90.5))
	logger.Error("Failed to open file", zap.String("file", "example.txt"), zap.Error(err))
	logger.Fatal("Database connection failed", zap.String("host", "localhost"), zap.Int("port", 3306))
}
