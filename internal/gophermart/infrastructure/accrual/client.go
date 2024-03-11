package accrual

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/psfpro/gophermart/internal/gophermart/domain"
	"io"
	"log"
	"net/http"
	"time"
)

type OrderProcessingResponse struct {
	Order   string
	Status  string
	Accrual float32
}

type Client struct {
	serverAddress         string
	transactionRepository domain.TransactionRepository
}

func NewClient(serverAddress string, transactionRepository domain.TransactionRepository) *Client {
	return &Client{serverAddress: serverAddress, transactionRepository: transactionRepository}
}

func (c *Client) ProcessOrder(ctx context.Context, order *domain.Transaction) error {
	number := order.OrderNumber()
	urlString := fmt.Sprintf("%s/api/orders/%s", c.serverAddress, number)
	request, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("accrual for order: %v response: %v", number, resp.Status)
	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return err
		}
		log.Printf("accrual body for order: %v response: %v", number, body)
		var data OrderProcessingResponse

		if err = json.Unmarshal(body, &data); err != nil {
			return err
		}
		if data.Status == "INVALID" || data.Status == "PROCESSED" {
			order.Processed(domain.TransactionStatus(data.Status), domain.NewTransactionAmount(data.Accrual))
			if err := c.transactionRepository.Save(ctx, order); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Client) Worker() {
	ctx := context.Background()
	for {
		log.Printf("accrual worker run")
		res, err := c.transactionRepository.GetNewAccruals(ctx)
		if err != nil {
			log.Printf("accrual get new accruals error: %v", err)
		}
		for _, v := range res {
			err := c.ProcessOrder(ctx, v)
			if err != nil {
				log.Printf("accrual process accrual error: %v", err)
			}
			time.Sleep(1 * time.Second)
		}
		time.Sleep(1 * time.Second)
	}
}
