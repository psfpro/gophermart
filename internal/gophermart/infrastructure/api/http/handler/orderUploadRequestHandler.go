package handler

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/psfpro/gophermart/internal/gophermart/application"
	"github.com/psfpro/gophermart/internal/gophermart/domain"
)

type OrderUploadRequestHandler struct {
	transactionService    *application.TransactionService
	authenticationService application.AuthenticationService
}

func NewOrderUploadRequestHandler(
	transactionService *application.TransactionService,
	authenticationService application.AuthenticationService,
) *OrderUploadRequestHandler {
	return &OrderUploadRequestHandler{
		transactionService:    transactionService,
		authenticationService: authenticationService,
	}
}

func (h *OrderUploadRequestHandler) HandleRequest(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	tokenString := request.Header.Get("Authorization")
	userID, err := h.authenticationService.GetUserID(tokenString)
	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		return
	}
	orderNumber, err := io.ReadAll(request.Body)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	num, err := strconv.Atoi(string(orderNumber))
	if err != nil || !ValidLuhn(num) {
		response.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	err = h.transactionService.NewAccrual(request.Context(), domain.NewUserID(userID), domain.OrderNumber(orderNumber))

	if err != nil {
		log.Println(err)
		if errors.Is(err, application.ErrTransactionUploadedByUser) {
			response.WriteHeader(http.StatusOK)
			return
		} else if errors.Is(err, application.ErrTransactionUploadedByOtherUser) {
			response.WriteHeader(http.StatusConflict)
			return
		}
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusAccepted)
}

// ValidLuhn check number is valid or not based on Luhn algorithm
func ValidLuhn(number int) bool {
	return (number%10+checksum(number/10))%10 == 0
}

func checksum(number int) int {
	var luhn int

	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 { // even
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		luhn += cur
		number = number / 10
	}
	return luhn % 10
}
