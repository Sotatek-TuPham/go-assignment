package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-server/config"
	"gin-server/dto"
	"gin-server/entity"
	"gin-server/util"
	"io/ioutil"
	"net/http"
	"time"
)

type PaymentService interface {
	CreatePayment(dto.CreatePayment) dto.PaymentResponse
}

type paymentService struct{}

func NewPaymentService() PaymentService {
	return &paymentService{}
}

func (service *paymentService) CreatePayment(orderParams dto.CreatePayment) dto.PaymentResponse {
	var payment entity.Payment
	payment.OrderID = orderParams.OrderId
	payment.Status = "CREATED"
	config.DB.Create(&payment)

	if orderParams.PIN != "" {
		time.AfterFunc(10*time.Second, func() {
			payload := dto.PaymentPayload{
				OrderID: orderParams.OrderId,
				Status:  util.RandomPaymentStatus(),
			}

			jsonData, err := json.Marshal(payload)
			if err != nil {
				panic(err)
			}

			url := "http://localhost:8080/orders/payment-hook"

			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
			if err != nil {
				panic(err)
			}
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(body))
		})
	}

	return dto.PaymentResponse{
		Message: "Payment Created",
	}
}
