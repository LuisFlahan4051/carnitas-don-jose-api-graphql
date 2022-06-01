package database

import (
	"context"
	"log"

	"github.com/TwiN/go-color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
}

const (
	DEFAULTPORT = "27017"
	DEFAULTHOST = "localhost"
)

func catch(err error) {
	if err != nil {
		log.Println(color.Ize(color.Red, err.Error()))
	}
}

func Connect(port string, host string) *Database {
	if port == "" {
		port = DEFAULTPORT
	}
	if host == "" {
		host = DEFAULTHOST
	}
	URI := "mongodb://" + host + ":" + port

	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	catch(err)

	err = client.Connect(context.TODO())
	catch(err)
	if err == nil {
		log.Println(color.Ize(color.Green, ">>> Connect to MongoDB Succesfully in :"+URI))
	}

	return &Database{
		Client: client,
	}
}

func TestConnection(port string, host string) {
	if port == "" {
		port = DEFAULTPORT
	}
	if host == "" {
		host = DEFAULTHOST
	}
	URI := "mongodb://" + host + ":" + port
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	catch(err)
	err = client.Connect(context.TODO())
	catch(err)
	err = client.Ping(context.TODO(), nil)
	catch(err)
	if err == nil {
		log.Println(color.Ize(color.Green, ">>> Connect to MongoDB Succesfully in "+URI))
	}
	err = client.Disconnect(context.TODO())
	catch(err)
}
