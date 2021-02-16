package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Data struct {
	db *sql.DB
}

type Response struct {
	Result string `json:"result"`
	Data interface{} `json:"data"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func NewData(db *sql.DB) *Data {
	return &Data{db: db}
}

func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Origin, X-Csrftoken, Content-Type, Accept")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
}

func ResponseByJSON(w http.ResponseWriter, code int, data interface{}) {
	SetHeaders(w)
	w.WriteHeader(code)
	response := Response{
		Result: http.StatusText(code),
		Data: data,
	}
	json.NewEncoder(w).Encode(response)
	return
}



func (data *Data) UserCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseByJSON(w, http.StatusMethodNotAllowed, nil)
		return
	}
}

func (data *Data) UserGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseByJSON(w, http.StatusMethodNotAllowed, nil)
		return
	}
}

func (data *Data) UserUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		ResponseByJSON(w, http.StatusMethodNotAllowed, nil)
		return
	}
}