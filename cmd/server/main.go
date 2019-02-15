package main

import (
	"log"
	"os"

	monitor360HTTP "github.com/nylo-andry/monitor360/http"
)

func main() {
	env := os.Getenv("ENV")
	allowedHost := os.Getenv("ALLOWED_HOST")
	if allowedHost == "" && env == "production" {
		log.Fatal("No ALLOWED_HOST provided")
	}

	log.Printf("ENV is %v", allowedHost)

	if env == "production" {
		monitor360HTTP.StartProductionServer(allowedHost)
	} else {
		monitor360HTTP.StartDevServer()
	}
}
