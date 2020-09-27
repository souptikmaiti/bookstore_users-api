package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(msg string) *RestErr {
	return &RestErr{msg, http.StatusBadRequest, "bad_request"}
}

func NewNotFoundError(msg string) *RestErr {
	return &RestErr{msg, http.StatusNotFound, "not_found"}
}
