package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHandlerRoutes(r *mux.Router) {
	r.Name("GetHelmReleases").Path("/api/v1/helm/releases").Methods("GET").HandlerFunc(GetHelmReleases)
	r.Name("GetHelmReleasesByNamespace").Path("/api/v1/helm/{cluster}/namespaces").Methods("GET").HandlerFunc(GetHelmReleasesByNamespace)
	r.Name("GetCurrentContext").Path("/api/v1/kubernetes/currentContext").Methods("GET").HandlerFunc(GetCurrentContext)
}

func JSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
