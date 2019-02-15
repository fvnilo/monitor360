package http

import (
	"log"
	"net/http"
)

func StartDevServer() {
	http.HandleFunc("/", handler)

	log.Printf("start listening at :http")
	log.Fatal(http.ListenAndServe(":http", nil))
}
