package usecases

import (
	"enigmacamp.com/coffeeshop/models"
	"enigmacamp.com/coffeeshop/repositories"
)

type TakeOrder interface {
	CreateCustomerOrder(orders []models.Coffee) *models.Order
}
type takeOrder struct {
	orderList repositories.OrderList
}

func (t *takeOrder) CreateCustomerOrder(orders []models.Coffee) *models.Order {
	orderId := "123"
	newOrder := models.NewOrder(orderId, orders)
	t.orderList.Create(&newOrder)
	return &newOrder
}

func NewTakeOrder(orderList repositories.OrderList) TakeOrder {
	takeOrder := takeOrder{
		orderList: orderList,
	}
	return &takeOrder
}
