package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler handler that returns HTML from a file.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("..", "index.html")
	http.ServeFile(w, r, path)

}

// UploadHandler parses the form, reads a file from it, passes this data for processing,
// creates a local file and writes the result to it.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error receiving form", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "File read error", http.StatusInternalServerError)
		return
	}

	content := string(data)

	converted, err := service.Detect(content)
	if err != nil {
		http.Error(w, "Processing error", http.StatusInternalServerError)
		return
	}

	extension := filepath.Ext(header.Filename)
	filename := time.Now().UTC().Format("20060102_150405") + extension
	out, err := os.Create(filename)
	if err != nil {
		http.Error(w, "File creation error", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = out.WriteString(converted)
	if err != nil {
		http.Error(w, "Recording error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(converted))
}
