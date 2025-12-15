package main

import (
	"net/http"
	"time"
)

var (
	backendConfigured bool
	backendReachable  bool
)

func setupBackend() {
	if config.Backend.URL != "" {
		backendConfigured = true

		go func() {
			ticker := time.NewTicker(time.Duration(config.Backend.CheckInterval) * time.Second)
			for range ticker.C {
				checkBackendStatus()
			}
		}()
	}
}

func checkBackendStatus() {
	resp, err := http.Get(config.Backend.URL)
	backendReachable = err == nil && resp.StatusCode == http.StatusOK
}
