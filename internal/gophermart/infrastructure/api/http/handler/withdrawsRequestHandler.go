package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/psfpro/gophermart/internal/gophermart/application"
)

type WithdrawalsRequestHandler struct {
	transactionService    *application.TransactionService
	authenticationService application.AuthenticationService
}

func NewWithdrawalsRequestHandler(
	transactionService *application.TransactionService,
	authenticationService application.AuthenticationService,
) *WithdrawalsRequestHandler {
	return &WithdrawalsRequestHandler{
		transactionService:    transactionService,
		authenticationService: authenticationService,
	}
}

func (h *WithdrawalsRequestHandler) HandleRequest(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	tokenString := request.Header.Get("Authorization")
	userID, err := h.authenticationService.GetUserID(tokenString)
	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		return
	}

	res, err := h.transactionService.GetWithdrawals(request.Context(), userID)

	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	type item struct {
		Order       string    `json:"order"`
		Sum         float32   `json:"sum"`
		ProcessedAt time.Time `json:"processed_at"`
	}
	var data []item

	for _, v := range res {
		data = append(data, item{
			Order:       string(v.OrderNumber()),
			Sum:         v.Amount().Value(),
			ProcessedAt: v.UpdatedAt(),
		})
	}

	jsonData, err := json.Marshal(data)
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
