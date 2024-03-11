package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/psfpro/gophermart/internal/gophermart/application"
	"github.com/psfpro/gophermart/internal/gophermart/domain"
)

type WithdrawRequest struct {
	Order string  `json:"order"`
	Sum   float32 `json:"sum"`
}

type WithdrawRequestHandler struct {
	transactionService    *application.TransactionService
	authenticationService application.AuthenticationService
}

func NewWithdrawRequestHandler(
	transactionService *application.TransactionService,
	authenticationService application.AuthenticationService,
) *WithdrawRequestHandler {
	return &WithdrawRequestHandler{
		transactionService:    transactionService,
		authenticationService: authenticationService,
	}
}

func (h *WithdrawRequestHandler) HandleRequest(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	tokenString := request.Header.Get("Authorization")
	userID, err := h.authenticationService.GetUserID(tokenString)
	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		return
	}
	var v WithdrawRequest
	response.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(request.Body).Decode(&v); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.transactionService.NewWithdrawal(
		request.Context(),
		domain.NewUserID(userID),
		domain.OrderNumber(v.Order),
		domain.NewTransactionAmount(v.Sum),
	)

	if err != nil {
		log.Println(err)
		if errors.Is(err, application.ErrUserUnauthorized) {
			response.WriteHeader(http.StatusUnauthorized)
			return
		}
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
}
