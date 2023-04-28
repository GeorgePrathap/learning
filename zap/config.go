package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Create a configuration for Zap logger
	cfg := zap.Config{
		Encoding:         "json",                               // Log encoding format (json, console)
		OutputPaths:      []string{"stdout", ".log"},           // Output destinations (stdout, file path)
		ErrorOutputPaths: []string{"stderr", ".errorLog"},      // Error output destinations (stderr, file path)
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel), // Log level (debug, info, warn, error)
		EncoderConfig: zapcore.EncoderConfig{ // Encoder configuration
			MessageKey:     "msg",                         // Log message field key
			LevelKey:       "level",                       // Log level field key
			TimeKey:        "time",                        // Log timestamp field key
			EncodeLevel:    zapcore.CapitalLevelEncoder,   // Log level capitalization (capital, lowercase, uppercase)
			EncodeTime:     zapcore.ISO8601TimeEncoder,    // Log timestamp format (ISO8601, EpochMillis, EpochNanos)
			EncodeDuration: zapcore.StringDurationEncoder, // Log duration format (string, nanos, millis, seconds)
		},
	}

	// Create a Zap logger from the configuration
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // Flushes any buffered log entries before the program exits

	// Use the logger to log messages
	logger.Debug("This is a debug message", zap.String("key", "value"))
	logger.Info("This is an info message", zap.Int("count", 42))
	logger.Warn("This is a warning message", zap.Float64("pi", 3.14))
	logger.Error("This is an error message", zap.Error(err))
}
