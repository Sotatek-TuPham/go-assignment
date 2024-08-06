package service

import (
	"fmt"
	"gin-server/config"
	"gin-server/dto"
	"gin-server/entity"
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
			resp, err := http.Get("http://localhost:8080/test")
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
