package main

import (
	"log"
	"net/http"

	"github.com/heptalium/httputil"
)

var alive bool = true

func aliveMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if alive {
			next.ServeHTTP(w, r)
		} else {
			httputil.WriteHttpStatus(w, http.StatusServiceUnavailable)
		}
	})
}

func handleLivez(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(http.StatusText(http.StatusOK) + "\n"))
}

func handleSetAlive(w http.ResponseWriter, r *http.Request) {
	alive = true
	log.Println("Service is now healthy.")
	w.Write([]byte("Service is now healthy\n"))
}

func handleSetDead(w http.ResponseWriter, r *http.Request) {
	alive = false
	log.Println("Service is now unhealthy.")
	w.Write([]byte("Service is now unhealthy\n"))
}
