package main

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type User struct {
	Name     string
	Age      int
	Location string
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	logger.Info("User logged in",
		zap.String("username", "john_doe"),
		zap.String("ip_address", "192.168.1.100"),
		zap.Time("timestamp", time.Now()),
	)

	logger.Info("Order processed",
		zap.String("order_id", "12345"),
		zap.Float64("total_amount", 99.99),
		zap.Bool("is_success", true),
		zap.Duration("processing_time", time.Millisecond*250),
	)

	user := User{
		Name:     "Alice",
		Age:      30,
		Location: "New York",
	}

	logger.Info("User profile",
		zap.Object("user", user),
	)
}

func (u User) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	// Implement serialization logic for User type
	// Use enc to encode fields of User type into the logging format
	enc.AddString("Name", u.Name)
	enc.AddInt("Age", u.Age)
	// add other fields as needed

	return nil
}
