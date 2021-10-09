package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func GetDatabase() *mongo.Database {
	return getClient().Database("instagram")
}

// Using singleton pattern for connecting to MongoDB
func getClient() *mongo.Client {
	if client != nil {
		return client
	}
	client = connect()
	return client
}

func connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI(db_url)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalln(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Connected to Database")
	return client
}

func Close(context context.Context) {
	if client == nil {
		return
	}
	err := client.Disconnect(context)
	if err != nil {
		log.Fatalln(err)
	}
}
