// package main

// import (
// 	"go.uber.org/zap"
// )

// func main() {
// 	logger, _ := zap.NewDevelopment()
// 	defer logger.Sync()

// 	logger.Info("Hello, world!")
// 	logger.Warn("This is a warning!")
// 	logger.Error("Oops, an error occurred.")
// }

// need to check again

package main

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Define a custom encoder function that formats log entries as key-value pairs
	keyValueEncoder := func(entry zapcore.Entry, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(fmt.Sprintf("%s=%s", "level", entry.Level.String()))
		encoder.AppendString(fmt.Sprintf("%s=%s", "time", entry.Time.Format(time.RFC3339)))
		encoder.AppendString(fmt.Sprintf("%s=%s", "message", entry.Message))
	}

	// Define a new core that uses the custom encoder for output
	customCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), // Use console encoder for console output
		zapcore.Lock(os.Stdout), // Lock the output stream so only one log entry can be written at a time
		zap.InfoLevel,           // Only log entries with info level or higher
	).With([]zapcore.Field{
		zapcore.Field{
			Key:    "message",
			Type:   zapcore.StringType,
			String: "message",
		},
		zapcore.Field{
			Key:    "time",
			Type:   zapcore.StringType,
			String: "time",
		},
		zapcore.Field{
			Key:    "level",
			Type:   zapcore.StringType,
			String: "level",
		},
		zapcore.Field{
			Key:     "caller",
			Type:    zapcore.StringType,
			Encoder: zapcore.ShortCallerEncoder,
		},
		zapcore.Field{
			Key:     "custom",
			Type:    zapcore.StringType,
			Encoder: keyValueEncoder,
		},
	})

	// Create a new logger instance that uses the custom core
	logger := zap.New(customCore)

	// Log some messages using the custom logger
	logger.Info("Hello, world!")
	logger.Warn("This is a warning!")
	logger.Error("Oops, an error occurred.")

	// Sync the logger to flush any buffered log entries
	logger.Sync()
}
