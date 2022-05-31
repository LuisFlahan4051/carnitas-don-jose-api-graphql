package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
}

const (
	PORT = "27017"
	URI  = "mongodb://localhost:" + PORT
)

func Connect() *Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		log.Println(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Println(err)
	}
	log.Println(">>> Connect to MongoDB Succesfully in " + URI)

	return &Database{
		Client: client,
	}
}

func TestConnection() {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		log.Println(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Println(err)
	}
	log.Println(">>> Connect to MongoDB Succesfully in " + URI)
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Println(err)
	}
}
