package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"cloud-storage-backend/handlers"

	_ "github.com/lib/pq"
)

const (
	dbUser     = "wlad"
	dbPassword = "Polde-7499"
	dbName     = "wlad"
)

var db *sql.DB

func main() {
	var err error
	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")

	handlers.SetDB(db)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Cloud Storage!")
	})
	http.HandleFunc("/upload", handlers.UploadHandler)
	http.HandleFunc("/download", handlers.DownloadHandler)
	http.HandleFunc("/files", handlers.FilesHandler)
	http.HandleFunc("/delete", handlers.DeleteHandler)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
