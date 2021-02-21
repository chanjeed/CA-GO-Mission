package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

type Data struct {
	db *sql.DB
}

type Response struct {
	Result string      `json:"result"`
	Data   interface{} `json:"data"`
}

type UserName struct {
	Name string `json:"name"`
}

type UserToken struct {
	Token string `json:"token"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func NewData(db *sql.DB) *Data {
	return &Data{db: db}
}

func GenerateToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Origin, X-Csrftoken, Content-Type, Accept")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func ResponseByJSON(w http.ResponseWriter, code int, data interface{}) {
	SetHeaders(w)
	w.WriteHeader(code)
	response := Response{
		Result: http.StatusText(code),
		Data:   data,
	}
	json.NewEncoder(w).Encode(response)
	return
}

func (data *Data) UserCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseByJSON(w, http.StatusMethodNotAllowed, nil)
		return
	}

	var userName UserName
	err := json.NewDecoder(r.Body).Decode(&userName)
	if err != nil {
		ResponseByJSON(w, http.StatusBadRequest, ErrorMessage{Message: "fail to decode request"})
		return
	}
	if userName.Name == "" {
		ResponseByJSON(w, http.StatusBadRequest, ErrorMessage{Message: "name is required"})
		return
	}
	token := GenerateToken(10)
	data.CreateUser(userName.Name, token)

	SetHeaders(w)
	w.WriteHeader(http.StatusOK)
	response := UserToken{
		Token: token,
	}
	json.NewEncoder(w).Encode(response)
	return
}

func (data *Data) UserGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseByJSON(w, http.StatusMethodNotAllowed, nil)
		return
	}

	userToken := r.Header.Get("x-token")

	if userToken == "" {
		ResponseByJSON(w, http.StatusBadRequest, ErrorMessage{Message: "token is required"})
		return
	}

	user, err := data.GetUserName(userToken)
	if err != nil {
		ResponseByJSON(w, http.StatusInternalServerError, ErrorMessage{Message: err.Error()})
		return
	}

	SetHeaders(w)
	w.WriteHeader(http.StatusOK)
	response := UserName{
		Name: user.Name,
	}
	json.NewEncoder(w).Encode(response)
	return
}

func (data *Data) UserUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		ResponseByJSON(w, http.StatusMethodNotAllowed, nil)
		return
	}
	userToken := r.Header.Get("x-token")
	var userName UserName
	err := json.NewDecoder(r.Body).Decode(&userName)
	if err != nil {
		ResponseByJSON(w, http.StatusBadRequest, ErrorMessage{Message: "fail to decode request"})
		return
	}
	if userName.Name == "" {
		ResponseByJSON(w, http.StatusBadRequest, ErrorMessage{Message: "name is required"})
		return
	}

	data.UpdateUser(userName.Name, userToken)

	ResponseByJSON(w, http.StatusOK, nil)
	return
}
