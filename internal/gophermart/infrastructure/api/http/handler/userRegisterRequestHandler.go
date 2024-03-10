package handler

import (
	"encoding/json"
	"github.com/psfpro/gophermart/internal/gophermart/application"
	"log"
	"net/http"
)

type UserRegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserRegisterRequestHandler struct {
	userService *application.UserService
}

func NewUserRegisterRequestHandler(userService *application.UserService) *UserRegisterRequestHandler {
	return &UserRegisterRequestHandler{userService: userService}
}

func (h *UserRegisterRequestHandler) HandleRequest(response http.ResponseWriter, request *http.Request) {
	var v UserRegisterRequest
	response.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(request.Body).Decode(&v); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.userService.Registration(request.Context(), v.Login, v.Password)

	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Authorization", res.AccessToken)
	response.WriteHeader(http.StatusOK)
}
