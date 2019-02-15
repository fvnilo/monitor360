package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nylo-andry/monitor360"
)

func getCrmStatus(w http.ResponseWriter) {
	url := "https://crmpro.sm360.ca"
	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)

	var r monitor360.MonitorResponse
	if err == nil && res.StatusCode == http.StatusOK {
		r = monitor360.NewResponse(false, "The CRM Pro is up!")
	} else {
		log.Printf("Got an error while pinging the CRM Pro: %v", err.Error())
		r = monitor360.NewResponse(false, "Uh oh! The CRM Pro is down!!!")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(r)
}
