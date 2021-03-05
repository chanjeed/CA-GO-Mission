package database

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
