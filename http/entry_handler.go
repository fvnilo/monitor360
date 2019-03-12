package http

import (
	"encoding/json"
	"net/http"

	"github.com/nylo-andry/monitor360"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	req := monitor360.MonitorRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()

	if err == nil && req.Result.Intent.DisplayName == "crm-status" {
		getCrmStatus(w)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Could not determine action."))
}
