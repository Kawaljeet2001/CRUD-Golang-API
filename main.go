package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Kawaljeet2001/netflix-api/database"
	"github.com/Kawaljeet2001/netflix-api/router"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("This is a proper Backend API")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")

	}
	port := os.Getenv("PORT")
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	go func() {
		database.DBInstance(context.Background())
	}()

	// defer func() {
	// 	cancel()
	// }()

	//register the routes here
	myrouter := router.CreateRouter()

	//connect to the port here
	log.Println("The server is up and runnning at port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, myrouter))
}
