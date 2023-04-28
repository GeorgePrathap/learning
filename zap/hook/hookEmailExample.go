package main

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/gomail.v2"
)

func main() {
	// Create a new logger instance
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	// Register a hook function that sends an email if the log level is Error or above
	logger = logger.WithOptions(zap.Hooks(emailAlertHook))

	// Log a message with level Info
	logger.Info("This is an info message")

	// Log a message with level Error
	logger.Error("This is an error message")

	// Sync the logger to ensure all logs are written
	logger.Sync()
}

// emailAlertHook is a hook function that sends an email if the log level is Error or above
func emailAlertHook(entry zapcore.Entry) error {
	if entry.Level >= zap.ErrorLevel {
		sendEmailAlert(entry.Message)
	}
	return nil
}

// sendEmailAlert sends an email alert
func sendEmailAlert(message string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "sender@example.com")
	m.SetHeader("To", "recipient@example.com")
	m.SetHeader("Subject", "Error Alert")
	m.SetBody("text/plain", message)

	d := gomail.NewDialer("smtp.example.com", 587, "username", "password")

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Failed to send email alert:", err)
	}
}
