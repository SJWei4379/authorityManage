package utils

import "net/http"

type Response struct {
	Code    int         `json:"code" bson:"code"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
}

func SuccessMess(message string, data interface{}) Response {
	return Response{
		http.StatusOK,
		message,
		data,
	}
}

func ErrorMess(message string, data interface{}) Response {
	return Response{
		http.StatusInternalServerError,
		message,
		data,
	}
}
