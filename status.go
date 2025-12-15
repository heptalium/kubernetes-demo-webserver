package main

import (
	"net/http"

	"github.com/heptalium/httputil"
	"gopkg.in/yaml.v2"
)

type backendStatus struct {
	Configured bool
	Reachable  bool
}

type status struct {
	Alive   bool
	Healthy bool
	Backend backendStatus
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	response := status{
		Alive:   alive,
		Healthy: healthy,
		Backend: backendStatus{
			Configured: backendConfigured,
			Reachable:  backendReachable,
		},
	}

	err := yaml.NewEncoder(w).Encode(response)
	if err != nil {
		httputil.WriteHttpStatus(w, http.StatusInternalServerError)
	}
}
