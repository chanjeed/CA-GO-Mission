package database

import (
	"database/sql"
	"fmt"
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
	DB = db
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func GetUserName(userToken string) (*User, error) {
	const sqlStr = `SELECT name FROM Users WHERE token=?`

	var name string

	var user User
	err := DB.QueryRow(sqlStr, userToken).Scan(&user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user with name %d", name)
		}
		return nil, err
	}

	return &user, nil
}

func CreateUser(userName string, userToken string) {
	const sqlStr = `INSERT INTO Users (name,token) VALUES (?,?);`

	ins, err := DB.Prepare(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	ins.Exec(userName, userToken)

	return
}

func UpdateUser(userName string, userToken string) {
	const sqlStr = `UPDATE Users SET name= ? WHERE token=?;`

	upd, err := DB.Prepare(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	upd.Exec(userName, userToken)

	return
}
