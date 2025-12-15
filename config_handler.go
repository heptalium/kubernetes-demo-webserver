package main

import (
	"net/http"

	"github.com/heptalium/httputil"
	"gopkg.in/yaml.v2"
)

func handleConfig(w http.ResponseWriter, r *http.Request) {
	err := yaml.NewEncoder(w).Encode(config)
	if err != nil {
		httputil.WriteHttpStatus(w, http.StatusInternalServerError)
	}
}
