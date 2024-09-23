package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "File ID is required", http.StatusBadRequest)
		return
	}

	db := GetDB()
	var filename string
	var fileData []byte
	err := db.QueryRow("SELECT filename, data FROM files WHERE id = $1", id).Scan(&filename, &fileData)
	if err != nil {
		log.Printf("Error retrieving file from database: %v", err)
		http.Error(w, "Unable to retrieve file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	w.Header().Set("Content-Type", "application/octet-stream")

	_, err = w.Write(fileData)
	if err != nil {
		log.Printf("Error writing file to response: %v", err)
		http.Error(w, "Unable to send file", http.StatusInternalServerError)
		return
	}

	log.Printf("File downloaded successfully: %s", filename)
}
