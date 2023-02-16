package usecases

import (
	"enigmacamp.com/coffeeshop/models"
	"enigmacamp.com/coffeeshop/repositories"
	"errors"
)

type Bill interface {
	CreateBill(orderId string) error
	GetListUnpaid() []models.Order
}
type bill struct {
	orderList repositories.OrderList
}

func (t *bill) CreateBill(orderId string) error {
	o, err := t.orderList.FindById(orderId)
	if err != nil {
		return errors.New("Order not found")
	}
	o.IsPaid = true
	return nil
}
func (t *bill) GetListUnpaid() []models.Order {
	unpaidOrder, err := t.orderList.FindByPaidStatus(false)
	if err != nil {
		return nil
	}
	return unpaidOrder
}

func NewBill(orderList repositories.OrderList) Bill {
	return &bill{
		orderList,
	}
}
