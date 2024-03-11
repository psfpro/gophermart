package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/psfpro/gophermart/internal/gophermart/application"
)

type OrdersRequestHandler struct {
	transactionService    *application.TransactionService
	authenticationService application.AuthenticationService
}

func NewOrdersRequestHandler(
	transactionService *application.TransactionService,
	authenticationService application.AuthenticationService,
) *OrdersRequestHandler {
	return &OrdersRequestHandler{
		transactionService:    transactionService,
		authenticationService: authenticationService,
	}
}

func (h *OrdersRequestHandler) HandleRequest(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	tokenString := request.Header.Get("Authorization")
	userID, err := h.authenticationService.GetUserID(tokenString)
	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		return
	}

	res, err := h.transactionService.GetAccruals(request.Context(), userID)

	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	type item struct {
		Number    string    `json:"number"`
		Status    string    `json:"status"`
		Accrual   float32   `json:"accrual"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	var data []item

	for _, v := range res {
		data = append(data, item{
			Number:    string(v.OrderNumber()),
			Status:    string(v.Status()),
			Accrual:   v.Amount().Value(),
			UpdatedAt: v.UpdatedAt(),
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
