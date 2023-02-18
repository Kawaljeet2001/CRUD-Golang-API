package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Kawaljeet2001/netflix-api/database"
	"github.com/Kawaljeet2001/netflix-api/router"
)

func main() {
	fmt.Println("This is a proper Backend API")
	// port := os.Getenv("PORT")
	port := 4000
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
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), myrouter))
}
