package main

import (
	"net/http"
	"os"
	"sort"
)

func handleEnvironment(w http.ResponseWriter, r *http.Request) {
	environ := os.Environ()
	sort.Strings(environ)

	for _, env := range environ {
		w.Write([]byte(env + "\n"))
	}
}
