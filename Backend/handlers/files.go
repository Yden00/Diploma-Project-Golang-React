package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func FilesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT id, filename FROM files")
	if err != nil {
		log.Printf("Error retrieving files from database: %v", err)
		http.Error(w, "Unable to retrieve files", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var files []map[string]interface{}
	for rows.Next() {
		var id int
		var filename string
		if err := rows.Scan(&id, &filename); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Error processing files", http.StatusInternalServerError)
			return
		}
		files = append(files, map[string]interface{}{
			"id":       id,
			"filename": filename,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(files); err != nil {
		log.Printf("Error encoding files to JSON: %v", err)
		http.Error(w, "Error encoding files", http.StatusInternalServerError)
		return
	}
}
