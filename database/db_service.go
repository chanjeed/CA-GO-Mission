package database

import (
	"database/sql"
	"fmt"
	"game/handler/user"
	"log"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func GetUserName(data *user.Data, userToken string) (*User, error) {
	const sqlStr = `SELECT name FROM Users WHERE token=?`

	var name string

	var user User
	err := data.DB.QueryRow(sqlStr, userToken).Scan(&user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user with name %d", name)
		}
		return nil, err
	}

	return &user, nil
}

func CreateUser(data *user.Data, userName string, userToken string) {
	const sqlStr = `INSERT INTO Users (name,token) VALUES (?,?);`

	ins, err := data.DB.Prepare(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	ins.Exec(userName, userToken)

	return
}

func UpdateUser(data *user.Data, userName string, userToken string) {
	const sqlStr = `UPDATE Users SET name= ? WHERE token=?;`

	upd, err := data.DB.Prepare(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	upd.Exec(userName, userToken)

	return
}
