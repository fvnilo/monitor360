package http

import (
	"crypto/tls"
	"log"
	"net/http"
	"sync"

	"golang.org/x/crypto/acme/autocert"
)

func StartProductionServer(allowedHost string) {
	mgr := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(allowedHost),
		Cache:      autocert.DirCache("certs"),
	}

	http.HandleFunc("/", handler)

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
