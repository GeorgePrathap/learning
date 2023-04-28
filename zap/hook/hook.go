package main

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Create a new logger instance
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// Register a hook function that sends an email if the log level is Error or above
	logger = logger.WithOptions(zap.Hooks(emailHook))

	// Register multiple hook functions
	// logger = logger.WithOptions(
	// 	zap.Hooks(
	// 		securityLogHook,
	// 		customLogHook,
	// 	),
	// )

	// Log a message with level Info
	logger.Info("This is an info message")

	// Log a message with level Error
	logger.Error("This is an error message")

	// Sync the logger to ensure all logs are written
	logger.Sync()
}

// emailHook is a hook function that sends an email if the log level is Error or above
func emailHook(entry zapcore.Entry) error {
	if entry.Level >= zap.ErrorLevel {
		sendEmail(entry.Message)
	}
	return nil
}

// sendEmail is a mock function that sends an email
func sendEmail(message string) {
	fmt.Println("Sending email:", message)
}
