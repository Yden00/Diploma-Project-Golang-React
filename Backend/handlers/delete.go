package handlers

import (
	"log"
	"net/http"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "File ID is required", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("DELETE FROM files WHERE id = $1", id)
	if err != nil {
		log.Printf("Error deleting file from database: %v", err)
		http.Error(w, "Unable to delete file", http.StatusInternalServerError)
		return
	}

	log.Printf("File deleted successfully: %s", id)
	w.WriteHeader(http.StatusNoContent)
}
