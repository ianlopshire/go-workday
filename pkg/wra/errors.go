package wra

import (
	"strconv"
)

type ResponseError struct {
	StatusCode int
	Message    string `json:"error"`
	Errors     []struct {
		Message  string `json:"error"`
		Location string `json:"location"`
	} `json:"errors"`
}

func (err ResponseError) Error() string {
	return "wra response error (" + strconv.Itoa(err.StatusCode) + "): " + err.Message
}
