package main

import (
	"fmt"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello World! (%s on %s)\n", variant, hostname)))
}
