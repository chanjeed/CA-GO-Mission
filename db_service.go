package main

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Token string `json:"token"`
}

type Character struct {
	ID int `json:"userCharacterID"`
	CharacterID int `json:"characterID"`
	Name string `json:"name"`
}

type Gacha struct {
	CharacterID int `json:"characterID"`
	Number int `json:"number"`
}

func (data *Data) GetUserName(userToken string) (*User, error) {
	const sqlStr = `SELECT name FROM Users WHERE token=?`
	
	var name string
	
	var user User
	err := data.db.QueryRow(sqlStr, userToken).Scan(&user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user with name %s", name)
		}
		return nil, err
	}

	return &user, nil
}

func (data *Data) CreateUser(userName string,userToken string) () {
	const sqlStr = `INSERT INTO Users (name,token) VALUES (?,?);`

	
	ins, err := data.db.Prepare(sqlStr)
    if err != nil {
        log.Fatal(err)
    }
    ins.Exec(userName,userToken)

	return
}

func (data *Data) UpdateUser(userName string,userToken string) () {
	const sqlStr = `UPDATE Users SET name= ? WHERE token=?;`

	
	upd, err := data.db.Prepare(sqlStr)
    if err != nil {
        log.Fatal(err)
    }
    upd.Exec(userName,userToken)

	return
}

func (data *Data) GetCharacterList(userToken string) ([]*Character,error) {
	const sqlStr = `SELECT c.*,d.name from (SELECT a.id,a.characterId from UserCharacters AS a JOIN Users AS b ON a.userId=b.id WHERE b.token=? ORDER BY a.ts DESC) AS c JOIN Characters AS d ON c.characterID=d.id;`
	

	
	var characterList []*Character
	rows, err := data.db.Query(sqlStr, userToken)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var character Character

		err := rows.Scan(&character.ID,&character.CharacterID,&character.Name)
		if err != nil {
			return nil, err
		}

		characterList = append(characterList, &character)
	}

	return characterList, nil
}

func (data *Data) GetGachaList() ([]*Gacha,error) {
	const sqlStr = `SELECT characterId,number from Gachas;`
		
	var GachaList []*Gacha
	rows, err := data.db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var gacha Gacha

		err := rows.Scan(&gacha.CharacterID,&gacha.Number)
		if err != nil {
			return nil, err
		}

		GachaList = append(GachaList, &gacha)
	}

	return GachaList, nil
}

func (data *Data) GetCharacterInfo() ([]*Gacha,error) {
	const sqlStr = `INSERT characterId,number from Gachas;`
		
	var GachaList []*Gacha
	rows, err := data.db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var gacha Gacha

		err := rows.Scan(&gacha.CharacterID,&gacha.Number)
		if err != nil {
			return nil, err
		}

		GachaList = append(GachaList, &gacha)
	}

	return GachaList, nil
}