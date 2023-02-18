package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/Kawaljeet2001/netflix-api/database"
	"github.com/Kawaljeet2001/netflix-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMovie(movie model.Netflix) {
	result, err := database.MoviesCollection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created a new movie with id: ", result.InsertedID)
}

func UpdateMovieById(movieId string) {
	//converting the movie id to primitive id
	mId, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": mId}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := database.MoviesCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count is: ", result.ModifiedCount)
}

func DeleteMovieById(movieId string) {
	mId, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": mId}
	result, err := database.MoviesCollection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The deleted count is: ", result.DeletedCount)
}

func DeleteAllMovies() {
	result, err := database.MoviesCollection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted all elements,: ", result.DeletedCount)
}

func GetAllMovies() []bson.M {
	cursor, err := database.MoviesCollection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []bson.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	//i want to return the movies

	return movies
}
