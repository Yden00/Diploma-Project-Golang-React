package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"cloud-storage-backend/handlers"
	"cloud-storage-backend/middleware"

	_ "github.com/lib/pq"
)

const (
	dbUser     = "wlad"
	dbPassword = "Polde-7499"
	dbName     = "wlad"
)

func main() {
	var err error
	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", psqlInfo)
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

	http.Handle("/", middleware.AuthMiddleware(http.HandlerFunc(welcomeHandler), false))
	http.Handle("/register", http.HandlerFunc(handlers.RegisterHandler))
	http.Handle("/login", http.HandlerFunc(handlers.LoginHandler))
	http.Handle("/upload", middleware.AuthMiddleware(http.HandlerFunc(handlers.UploadHandler), true))
	http.Handle("/download", middleware.AuthMiddleware(http.HandlerFunc(handlers.DownloadHandler), true))
	http.Handle("/files", middleware.AuthMiddleware(http.HandlerFunc(handlers.FilesHandler), true))
	http.Handle("/delete", middleware.AuthMiddleware(http.HandlerFunc(handlers.DeleteHandler), true))

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Cloud Storage!")
}
