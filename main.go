package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// DB の接続情報
const (
        DRIVER_NAME      = "mysql" // ドライバ名(mysql固定)
        // user:password@tcp(container-name:port)/dbname ※mysql はデフォルトで用意されているDB
        DATA_SOURCE_NAME = "root:mysql@tcp(127.0.0.1:3306)/game"
)


func main() {
	db, connectionError := sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
        if connectionError != nil {
                log.Fatal("error connecting to database: ", connectionError)
        }

	data := NewData(db)
	http.HandleFunc("/user/create", data.UserCreate)
	http.HandleFunc("/user/get", data.UserGet)
	http.HandleFunc("/user/update", data.UserUpdate)


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("The Server runs with http://localhost:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

