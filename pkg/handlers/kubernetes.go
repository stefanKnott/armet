package handlers

import (
	"net/http"

	k "github.com/stefanKnott/armet/pkg/kubernetes"
)

type getKubernetesContextResponse struct {
	Context string `json:"context"`
}

func GetCurrentContext(w http.ResponseWriter, r *http.Request) {
	getKubernetesContextResponse := getKubernetesContextResponse{
		Context: k.CurrentContext,
	}
	JSON(w, 200, getKubernetesContextResponse)
}
