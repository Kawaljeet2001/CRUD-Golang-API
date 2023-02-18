package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Kawaljeet2001/netflix-api/database"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("This is a proper Backend API")
	// port := os.Getenv("PORT")
	port := 4000
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	go func() {
		database.DBInstance(ctx)
	}()

	defer func() {
		cancel()
	}()

	//register the routes here
	router := mux.NewRouter()

	//connect to the port here
	log.Println("The server is up and runnning at port: ", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
