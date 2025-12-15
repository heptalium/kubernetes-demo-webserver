package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func configureLogging() {
	if config.LogFile != "" {
		log.Println("Writing logs to", config.LogFile)

		// Create log file directory if necessary
		err := os.MkdirAll(filepath.Dir(config.LogFile), 0777)
		if err != nil {
			log.Fatalln("Cannot create log file directory:", err)
		}

		// Open log file for writing in append mode
		file, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalln("Cannot open log file for writing:", err)
		}

		// Use log file for logging
		log.SetOutput(file)

		// Log startup message to log file
		logStartupMessage()
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func logRequest(r *http.Request, status int) {
	log.Println(r.RemoteAddr, r.Method, r.URL.Path, status)
}

func logStartupMessage() {
	log.Printf("Starting webserver %s on %s.", variant, hostname)
}
