package database

import (
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client;


func initDB() error {
	var err error
	Client, err = mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {

		log.Fatal(err)
		return err
	}
	return nil
}

func Handler() {
	err := initDB()

	if err != nil {
		log.Fatal(err)
	}
}