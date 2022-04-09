package helpers

import "strings"

type Response struct {
	Status bool        `json:"status"`
	Errors interface{} `json:"errors"`
	Data   interface{} `json:"data"`
}

func BuildResponse(data interface{}) Response {
	resp := Response{
		Status: true,
		Errors: nil,
		Data:   data,
	}

	return resp
}

func BuildErrResponse(err string) Response {
	splitErr := strings.Split(err, "\n")

	resp := Response{
		Status: false,
		Errors: splitErr,
		Data:   nil,
	}

	return resp
}
