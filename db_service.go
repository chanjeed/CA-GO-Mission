package main

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func (data *Data) GetUserName(userToken string) (*User, error) {
	const sqlStr = `SELECT name FROM Users WHERE token=?`

	var name string

	var user User
	err := data.db.QueryRow(sqlStr, userToken).Scan(&user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user with name %d", name)
		}
		return nil, err
	}

	return &user, nil
}

func (data *Data) CreateUser(userName string, userToken string) {
	const sqlStr = `INSERT INTO Users (name,token) VALUES (?,?);`

	ins, err := data.db.Prepare(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	ins.Exec(userName, userToken)

	return
}

func (data *Data) UpdateUser(userName string, userToken string) {
	const sqlStr = `UPDATE Users SET name= ? WHERE token=?;`

	upd, err := data.db.Prepare(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	upd.Exec(userName, userToken)

	return
}
