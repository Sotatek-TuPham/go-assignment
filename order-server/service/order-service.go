package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gin-server/config"
	"gin-server/dto"
	"gin-server/entity"
	utility "gin-server/util"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type OrderService interface {
	CreateOrder(dto.CreateOrder) entity.Order
	PaymentHook(dto.PaymentPayload) dto.OrderResponse
	CancelOrder(string) (entity.Order, error)
}

type orderService struct{}

func NewOrderService() OrderService {
	return &orderService{}
}

func (service *orderService) callToPaymentService(orderId uuid.UUID) {
	payload := dto.OrderPayload{
		OrderID: orderId,
		PIN:     utility.GeneratePIN(),
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	url := "http://localhost:8081/payments"

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
}

func (service *orderService) CreateOrder(orderParams dto.CreateOrder) entity.Order {
	var order entity.Order
	order.Email = orderParams.Email
	order.ID = uuid.New()
	order.Quantity = orderParams.Quantity
	order.UnitPrice = uint(rand.Intn(100))
	order.TotalPrice = int64(order.UnitPrice) * int64(order.Quantity)
	order.Status = entity.CREATED
	config.DB.Create(&order)

	service.callToPaymentService(order.ID)

	return order
}

func (service *orderService) PaymentHook(paymentPayload dto.PaymentPayload) dto.OrderResponse {
	var order entity.Order
	config.DB.Where("id = ?", paymentPayload.OrderID).First(&order)
	order.Status = entity.Status(paymentPayload.Status)
	config.DB.Save(&order)

	if entity.Status(paymentPayload.Status) == entity.CONFIRMED {
		time.AfterFunc(10*time.Second, func() {
			var order entity.Order
			config.DB.Where("id = ?", paymentPayload.OrderID).First(&order)

			// Update the order status to DELIVERED if it is still CONFIRMED
			if order.Status == entity.CONFIRMED {
				order.Status = entity.Status(entity.DELIVERED)
				config.DB.Save(&order)
			}
		})
	}
	return dto.OrderResponse{Message: "Order updated"}
}

func (service *orderService) CancelOrder(orderId string) (entity.Order, error) {
	var order entity.Order
	config.DB.Where("id = ?", orderId).First(&order)
	if order.Status == entity.CREATED {
		return order, errors.New("cannot update order in Created state")
	}

	if order.Status == entity.DELIVERED {
		return order, errors.New("cannot update order in Delivered state")
	}

	order.Status = entity.Status(entity.CANCELLED)
	config.DB.Save(&order)
	return order, nil
}
