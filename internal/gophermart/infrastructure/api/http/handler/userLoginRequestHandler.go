package handler

import (
	"encoding/json"
	"errors"
	"github.com/psfpro/gophermart/internal/gophermart/application"
	"log"
	"net/http"
)

type UserLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserLoginRequestHandler struct {
	userService *application.UserService
}

func NewUserLoginRequestHandler(userService *application.UserService) *UserLoginRequestHandler {
	return &UserLoginRequestHandler{userService: userService}
}

func (h *UserLoginRequestHandler) HandleRequest(response http.ResponseWriter, request *http.Request) {
	var v UserLoginRequest
	response.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(request.Body).Decode(&v); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.userService.Login(request.Context(), v.Login, v.Password)

	if err != nil {
		log.Println(err)
		if errors.Is(err, application.ErrUserUnauthorized) {
			response.WriteHeader(http.StatusUnauthorized)
			return
		}
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Authorization", res.AccessToken)
	response.WriteHeader(http.StatusOK)
}
