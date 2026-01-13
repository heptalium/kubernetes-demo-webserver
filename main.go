package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	config   Config
	hostname string
	variant  string = "black"
)

func init() {
	hostname, _ = os.Hostname()
}

func main() {
	logStartupMessage()

	// Load config from config file and environemnt variables
	configure()

	// Change log target if requested
	configureLogging()

	// Enable backend check if configured
	setupBackend()

	// Create data directory if configured
	setupDataDir()

	// Wait some time before serving requests
	go func() {
		time.Sleep(time.Duration(config.StartDelay) * time.Second)
		setHealthy()
	}()

	// Define all handlers for /<handler>, /api/<handler> and /<variant>/<handler>
	for _, prefix := range []string{"", "/api", "/" + variant} {
		http.Handle(prefix+"/", loggingMiddleware(healthyMiddleware(http.HandlerFunc(handleIndex))))

		http.Handle(prefix+"/healthz", healthyMiddleware(http.HandlerFunc(handleHealthz)))
		http.Handle(prefix+"/set-healthy", loggingMiddleware(http.HandlerFunc(handleSetHealthy)))
		http.Handle(prefix+"/set-unhealthy", loggingMiddleware(http.HandlerFunc(handleSetUnhealthy)))

		http.Handle(prefix+"/livez", aliveMiddleware(http.HandlerFunc(handleHealthz)))
		http.Handle(prefix+"/set-alive", loggingMiddleware(http.HandlerFunc(handleSetAlive)))
		http.Handle(prefix+"/set-dead", loggingMiddleware(http.HandlerFunc(handleSetDead)))

		http.Handle(prefix+"/crash", loggingMiddleware(http.HandlerFunc(handleCrash)))

		http.Handle(prefix+"/data/", loggingMiddleware(healthyMiddleware(http.StripPrefix(prefix+"/data", http.FileServer(http.Dir("/"))))))

		http.Handle(prefix+"/env", loggingMiddleware(healthyMiddleware(http.HandlerFunc(handleEnvironment))))
		http.Handle(prefix+"/environment", loggingMiddleware(healthyMiddleware(http.HandlerFunc(handleEnvironment))))

		http.Handle(prefix+"/config", loggingMiddleware(healthyMiddleware(http.HandlerFunc(handleConfig))))
		http.Handle(prefix+"/status", loggingMiddleware(http.HandlerFunc(handleStatus)))

		http.Handle(prefix+"/upload", loggingMiddleware(healthyMiddleware(http.HandlerFunc(handleUpload))))

		http.Handle(prefix+"/color", loggingMiddleware(healthyMiddleware(http.HandlerFunc(handleColor))))
	}

	// Start web service
	log.Printf("Listen on port :%d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}
