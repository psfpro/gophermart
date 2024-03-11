package handler

import (
	"encoding/json"
	"github.com/psfpro/gophermart/internal/gophermart/application"
	"log"
	"net/http"
)

type BalanceRequestHandler struct {
	transactionService    *application.TransactionService
	authenticationService application.AuthenticationService
}

func NewBalanceRequestHandler(
	transactionService *application.TransactionService,
	authenticationService application.AuthenticationService,
) *BalanceRequestHandler {
	return &BalanceRequestHandler{
		transactionService:    transactionService,
		authenticationService: authenticationService,
	}
}

func (h *BalanceRequestHandler) HandleRequest(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	tokenString := request.Header.Get("Authorization")
	userID, err := h.authenticationService.GetUserID(tokenString)
	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		return
	}

	account, err := h.transactionService.GetBalance(request.Context(), userID)
	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(struct {
		Current   float32 `json:"current"`
		Withdrawn float32 `json:"withdrawn"`
	}{
		Current:   account.Balance.Value(),
		Withdrawn: account.Withdrawals.Value(),
	})
	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	_, err = response.Write(jsonData)
	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
}
