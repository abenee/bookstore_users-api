package errors

import "net/http"

type ErrorResp struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *ErrorResp {
	return &ErrorResp{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
	}
}

func NewRecordNotfoundError(message string) *ErrorResp {
	return &ErrorResp{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Bad Request",
	}
}
