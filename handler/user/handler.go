package user

import (
	"encoding/json"
	"game/database"
	"game/util"
	"net/http"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
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
	token := util.GenerateToken(10)

	database.CreateUser(userName.Name, token)

	SetHeaders(w)
	w.WriteHeader(http.StatusOK)
	response := UserToken{
		Token: token,
	}
	json.NewEncoder(w).Encode(response)
	return
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseByJSON(w, http.StatusMethodNotAllowed, nil)
		return
	}

	userToken := r.Header.Get("x-token")

	if userToken == "" {
		ResponseByJSON(w, http.StatusBadRequest, ErrorMessage{Message: "token is required"})
		return
	}

	user, err := database.GetUserName(userToken)
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

func UserUpdate(w http.ResponseWriter, r *http.Request) {
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

	database.UpdateUser(userName.Name, userToken)

	ResponseByJSON(w, http.StatusOK, nil)
	return
}
