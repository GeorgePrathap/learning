package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Create a new logger instance
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	// Register a hook function that logs to a MongoDB database
	logger = logger.WithOptions(zap.Hooks(mongoDBLogHook))

	// Log a message
	logger.Info("This is a message")

	// Sync the logger to ensure all logs are written
	logger.Sync()
}

// mongoDBLogHook is a hook function that logs to a MongoDB database
func mongoDBLogHook(entry zapcore.Entry) error {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	defer client.Disconnect(context.Background())

	// Get a handle for the "logs" collection
	collection := client.Database("testdb").Collection("logs")

	// Insert the log entry into the "logs" collection
	_, err = collection.InsertOne(context.Background(), entry)
	if err != nil {
		return fmt.Errorf("failed to insert log into MongoDB: %w", err)
	}

	return nil
}
