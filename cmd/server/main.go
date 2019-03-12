package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	monitor360HTTP "github.com/nylo-andry/monitor360/http"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No PORT provided")
	}

	log.Printf("Starting server on port %v", port)

	http.HandleFunc("/", monitor360HTTP.MainHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
