package handlers

import (
	"net/http"
)

type getKubernetesContextResponse struct {
	Context string `json:"context"`
}

func GetCurrentContext(w http.ResponseWriter, r *http.Request) {
	getKubernetesContextResponse := getKubernetesContextResponse{
		Context: "",
	}
	JSON(w, 200, getKubernetesContextResponse)
}
