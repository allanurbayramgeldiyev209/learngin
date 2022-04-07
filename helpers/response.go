package helpers

import "strings"

type Response struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Errors interface{} `json:"errors"`
	Data interface{} `json:"data"`
}

func BuildResponse(message string, data interface{}) Response {
	resp := Response{
		Status: true,
		Message: message,
		Errors: nil,
		Data: data,
	}

	return resp
}

func BuildErrResponse(message string, err string,data interface{}) Response {
	splitErr := strings.Split(err,"\n")

	resp := Response{
		Status: false,
		Message: message,
		Errors: splitErr,
		Data: data,
	}

	return resp
}