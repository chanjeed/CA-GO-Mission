package user

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Result string      `json:"result"`
	Data   interface{} `json:"data"`
}

type ErrorMessage struct {
	Message string `json:"message"`
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
