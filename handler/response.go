package handler

type Response struct {
	Result string      `json:"result"`
	Data   interface{} `json:"data"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}
