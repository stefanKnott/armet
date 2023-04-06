package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stefanKnott/armet/pkg/helm"
)

type GetHelmReleasesResponse struct {
	Releases map[string]map[string][]helm.HelmRelease `json:"releases"`
}

type GetHelmReleasesByNamespaceResponse struct {
	Releases map[string][]helm.HelmRelease `json:"releases"`
}

func GetHelmReleases(w http.ResponseWriter, r *http.Request) {
	releases := helm.GetReleases()
	getHelmReleasesResponse := GetHelmReleasesResponse{
		Releases: releases,
	}
	JSON(w, 200, getHelmReleasesResponse)
}

func GetHelmReleasesByNamespace(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	releases := helm.GetReleaseByNamespaceMap(params["cluster"])
	getHelmReleasesByNamespaceResponse := GetHelmReleasesByNamespaceResponse{
		Releases: releases,
	}
	JSON(w, 200, getHelmReleasesByNamespaceResponse)
}
