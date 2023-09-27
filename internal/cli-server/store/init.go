package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	mongo *mongo.Database
}

var db *Database

func InitDB() *mongo.Client {
	// Create a top context
	// Setting timeout for the context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	m := client.Database("cli-server")
	database := Database{
		mongo: m,
	}
	db = &database
	return client
}
