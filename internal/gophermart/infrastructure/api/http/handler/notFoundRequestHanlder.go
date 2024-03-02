package handler

import (
	"net/http"
)

type NotFoundRequestHandler struct {
}

func NewNotFoundRequestHandler() *NotFoundRequestHandler {
	return &NotFoundRequestHandler{}
}

func (obj *NotFoundRequestHandler) HandleRequest(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusNotFound)
}
