package user

import (
	"database/sql"
	"encoding/json"
	"game/database"
	"game/handler"
	"game/util"
	"net/http"
)

type Data struct {
	DB *sql.DB
}

func NewData(db *sql.DB) *Data {
	return &Data{DB: db}
}

func (data *Data) UserCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handler.ResponseByJSON(w, http.StatusMethodNotAllowed, nil)
		return
	}

	var userName UserName
	err := json.NewDecoder(r.Body).Decode(&userName)
	if err != nil {
		handler.ResponseByJSON(w, http.StatusBadRequest, handler.ErrorMessage{Message: "fail to decode request"})
		return
	}
	if userName.Name == "" {
		handler.ResponseByJSON(w, http.StatusBadRequest, handler.ErrorMessage{Message: "name is required"})
		return
	}
	token := util.GenerateToken(10)
	database.CreateUser(data, userName.Name, token)

	handler.SetHeaders(w)
	w.WriteHeader(http.StatusOK)
	response := UserToken{
		Token: token,
	}
	json.NewEncoder(w).Encode(response)
	return
}

func (data *Data) UserGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handler.ResponseByJSON(w, http.StatusMethodNotAllowed, nil)
		return
	}

	userToken := r.Header.Get("x-token")

	if userToken == "" {
		handler.ResponseByJSON(w, http.StatusBadRequest, handler.ErrorMessage{Message: "token is required"})
		return
	}

	user, err := database.GetUserName(data, userToken)
	if err != nil {
		handler.ResponseByJSON(w, http.StatusInternalServerError, handler.ErrorMessage{Message: err.Error()})
		return
	}

	handler.SetHeaders(w)
	w.WriteHeader(http.StatusOK)
	response := UserName{
		Name: user.Name,
	}
	json.NewEncoder(w).Encode(response)
	return
}

func (data *Data) UserUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		handler.ResponseByJSON(w, http.StatusMethodNotAllowed, nil)
		return
	}
	userToken := r.Header.Get("x-token")
	var userName UserName
	err := json.NewDecoder(r.Body).Decode(&userName)
	if err != nil {
		handler.ResponseByJSON(w, http.StatusBadRequest, handler.ErrorMessage{Message: "fail to decode request"})
		return
	}
	if userName.Name == "" {
		handler.ResponseByJSON(w, http.StatusBadRequest, handler.ErrorMessage{Message: "name is required"})
		return
	}

	database.UpdateUser(data, userName.Name, userToken)

	handler.ResponseByJSON(w, http.StatusOK, nil)
	return
}
