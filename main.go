package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := sql.Open("sqlite3", "sample.db")
	if err != nil {
		log.Fatal(err)
	}
	data := NewData(db)

	http.HandleFunc("/users/get", data.UserHandler)
	http.HandleFunc("/items/get", data.ItemHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("The Server runs with http://localhost:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
