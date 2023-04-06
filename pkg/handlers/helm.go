package handlers

import (
	"net/http"

	"github.com/stefanKnott/armet/pkg/helm"
)

type GetHelmReleasesResponse struct {
	Releases []helm.HelmRelease `json:"releases"`
}

type GetHelmReleasesByNamespaceResponse struct {
	Releases map[string][]helm.HelmRelease `json:"releases"`
}

func GetHelmReleases(w http.ResponseWriter, r *http.Request) {
	releases := helm.GetReleaseSlice()
	getHelmReleasesResponse := GetHelmReleasesResponse{
		Releases: releases,
	}
	JSON(w, 200, getHelmReleasesResponse)
}

func GetHelmReleasesByNamespace(w http.ResponseWriter, r *http.Request) {
	releases := helm.GetReleaseByNamespaceMap()
	getHelmReleasesByNamespaceResponse := GetHelmReleasesByNamespaceResponse{
		Releases: releases,
	}
	JSON(w, 200, getHelmReleasesByNamespaceResponse)
}
