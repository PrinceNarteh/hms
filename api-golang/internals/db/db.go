// Package db
package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"hms-backend/internals/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectDB() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(config.Envs.DBConfig.URI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connected successfully")
	return client.Database(config.Envs.DBConfig.Name), nil
}
