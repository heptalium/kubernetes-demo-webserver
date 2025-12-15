package main

import (
	"log"
	"net/http"
)

func handleCrash(w http.ResponseWriter, r *http.Request) {
	log.Println("Service will crash now...")
	w.Write([]byte("Service will crash now...\n"))
	go func() {
		panic("Intentional crash triggered via /crash endpoint")
	}()
}
