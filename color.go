package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"

	"github.com/heptalium/httputil"
)

var colors = []string{"red", "orange", "yellow", "green", "blue", "purple"}

func handleColor(w http.ResponseWriter, r *http.Request) {
	if !backendReachable {
		httputil.WriteHttpStatus(w, http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		httputil.WriteHttpStatus(w, http.StatusMethodNotAllowed)
		return
	}

	field := r.URL.Query().Get("field")
	if field == "" {
		httputil.WriteHttpStatus(w, http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("%s/color%s", config.Backend.URL, field)

	// Get current color from backend
	resp, err := http.Get(url)
	if err != nil {
		httputil.WriteHttpStatus(w, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var color string

	// Determine current color, white if undefined
	switch resp.StatusCode {
	case http.StatusOK:
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			httputil.WriteHttpStatus(w, http.StatusInternalServerError)
			return
		}
		color = strings.Trim(string(body), "\n")
	case http.StatusNotFound:
		color = "white"
	default:
		httputil.WriteHttpStatus(w, http.StatusInternalServerError)
		return
	}

	// On POST requests select a new color which differs from the current one
	if r.Method == http.MethodPost {
		currentColor := color
		for {
			color = colors[rand.Intn(len(colors))]
			if color != currentColor {
				break
			}
		}

		// Save new color in the backend
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer([]byte(color)))
		if err != nil {
			httputil.WriteHttpStatus(w, http.StatusInternalServerError)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			httputil.WriteHttpStatus(w, http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
	}

	w.Write([]byte(fmt.Sprintf("%s\n", color)))
}
