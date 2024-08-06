package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-server/config"
	"gin-server/dto"
	"gin-server/entity"
	utility "gin-server/util"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/google/uuid"
)

type OrderService interface {
	CreateOrder(dto.CreateOrder) entity.Order
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
