package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var MoviesCollection *mongo.Collection

func DBInstance(ctx context.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")

	}
	mongodb_uri := os.Getenv("MONGODB_URI")
	if mongodb_uri == "" {
		log.Fatal("Error: not able to load MONGODB_URI")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongodb_uri))

	if err != nil {
		panic(err)
	}

	log.Println("Connected to Database Successfully!")

	Client = client
	MoviesCollection = client.Database("my-database").Collection("movies")
	// // <-ctx.Done()
	// if err := client.Disconnect(context.TODO()); err != nil {
	// 	panic(err)
	// }

	// log.Panicln("Disconnected MongoDB")

}
