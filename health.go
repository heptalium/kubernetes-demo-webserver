package main

import (
	"log"
	"net/http"

	"github.com/heptalium/httputil"
)

var healthy bool = false

func healthyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if healthy {
			next.ServeHTTP(w, r)
		} else {
			httputil.WriteHttpStatus(w, http.StatusServiceUnavailable)
		}
	})
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(http.StatusText(http.StatusOK) + "\n"))
}

func handleSetHealthy(w http.ResponseWriter, r *http.Request) {
	setHealthy()
	w.Write([]byte("Service is now healthy\n"))
}

func handleSetUnhealthy(w http.ResponseWriter, r *http.Request) {
	setUnhealthy()
	w.Write([]byte("Service is now unhealthy\n"))
}

func setHealthy() {
	healthy = true
	log.Println("Service is now healthy.")
}

func setUnhealthy() {
	healthy = false
	log.Println("Service is now unhealthy.")
}
