package main

import (
	"fmt"
	"game/database"
	"game/handler/user"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.ConnectToMySQL()
	http.HandleFunc("/user/create", user.UserCreate)
	http.HandleFunc("/user/get", user.UserGet)
	http.HandleFunc("/user/update", user.UserUpdate)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("The Server runs with http://localhost:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
