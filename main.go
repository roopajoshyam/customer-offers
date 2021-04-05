package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/jroopa/offers/routes"
)

func main() {
	e := godotenv.Load()

	if e != nil {
		fmt.Print("error loading env")
	}
	fmt.Println(e)
	port := os.Getenv("PORT")

	// Handle routes
	http.Handle("/", routes.Handlers())

	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
