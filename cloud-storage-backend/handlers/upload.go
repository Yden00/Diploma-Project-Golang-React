package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(1 << 30) // 1 GB
	if err != nil {
		log.Printf("Error parsing form: %v", err)
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error retrieving file: %v", err)
		http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	chunkSize := 1024 * 1024
	buf := make([]byte, chunkSize)

	var fileData []byte

	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			log.Printf("Error reading file: %v", err)
			http.Error(w, "Unable to read file", http.StatusInternalServerError)
			return
		}
		if n == 0 {
			break
		}
		fileData = append(fileData, buf[:n]...)
	}

	_, err = db.Exec("INSERT INTO files (filename, data) VALUES ($1, $2)", fileHeader.Filename, fileData)
	if err != nil {
		log.Printf("Error saving file to database: %v", err)
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	log.Printf("File uploaded successfully: %s", fileHeader.Filename)
	fmt.Fprintf(w, "File uploaded successfully!")
}
