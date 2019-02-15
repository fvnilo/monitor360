package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"sync"

	monitor360HTTP "github.com/nylo-andry/monitor360/http"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	env := os.Getenv("ENV")
	allowedHost := os.Getenv("ALLOWED_HOST")
	if allowedHost == "" && env == "production" {
		log.Fatal("No ALLOWED_HOST provided")
	}

	mgr := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(allowedHost),
		Cache:      autocert.DirCache("certs"),
	}

	http.HandleFunc("/", monitor360HTTP.Handler)

	var wg sync.WaitGroup

	server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: mgr.GetCertificate,
		},
	}

	wg.Add(2)

	go func() {
		log.Printf("start listening at :http")
		log.Fatal(http.ListenAndServe(":http", mgr.HTTPHandler(nil)))
		wg.Done()
	}()

	go func() {
		log.Printf("start listening at :https")
		log.Fatal(server.ListenAndServeTLS("", "")) // Key and cert provided by Let's Encrypt
		wg.Done()
	}()

	wg.Wait()

	log.Println("All servers have stopped.")
}
