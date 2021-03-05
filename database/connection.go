package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

// DB の接続情報
const (
	DriverName = "mysql" // ドライバ名(mysql固定)
	// user:password@tcp(container-name:port)/dbname ※mysql はデフォルトで用意されているDB
	DataSourceName = "root:mysql@tcp(127.0.0.1:3306)/game"
)

func ConnectToMySQL() {
	db, connectionError := sql.Open(DriverName, DataSourceName)
	if connectionError != nil {
		log.Fatal("error connecting to database: ", connectionError)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	DB = db
}
