package http

import (
	"encoding/json"
	"net/http"

	"github.com/nylo-andry/monitor360"
)

func handler(w http.ResponseWriter, r *http.Request) {
	req := monitor360.MonitorRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()

	if err == nil && req.Result.Intent.DisplayName == "crm-status" {
		getCrmStatus(w)
	}
}
