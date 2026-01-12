package main

import (
	"errors"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/heptalium/httputil"
)

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if config.DataDir == "" {
		httputil.WriteHttpStatus(w, http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		httputil.WriteHttpStatus(w, http.StatusMethodNotAllowed)
		return
	}

	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		httputil.WriteHttpStatus(w, http.StatusBadRequest)
		return
	}
	if mediaType != "multipart/form-data" {
		w.Header().Set("Accept", "multipart/form-data")
		httputil.WriteHttpStatus(w, http.StatusUnsupportedMediaType)
		return
	}

	r.ParseMultipartForm(8 << 20)

	upload, header, err := r.FormFile("file")
	if err != nil {
		httputil.WriteHttpStatus(w, http.StatusBadRequest)
		return
	}
	defer upload.Close()

	file, err := os.Create(filepath.Join(config.DataDir, header.Filename))
	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			httputil.WriteHttpStatus(w, http.StatusForbidden)
		} else {
			httputil.WriteHttpStatus(w, http.StatusInternalServerError)
		}
		return
	}
	defer file.Close()

	_, err = io.Copy(file, upload)
	if err != nil {
		httputil.WriteHttpStatus(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", filepath.Join("/data", config.DataDir, header.Filename))
	httputil.WriteHttpStatus(w, http.StatusCreated)
}
