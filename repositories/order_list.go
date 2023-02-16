package repositories

import (
	"enigmacamp.com/coffeeshop/models"
	"errors"
)

type OrderList interface {
	Create(order *models.Order)
	FindById(orderId string) (*models.Order, error)
	FindByPaidStatus(billStatus bool) ([]models.Order, error)
}
type orderList struct {
	orders []*models.Order
}

func (c *orderList) Create(order *models.Order) {
	c.orders = append(c.orders, order)
}

func (c *orderList) FindById(orderId string) (*models.Order, error) {
	for _, c := range c.orders {
		if c.OrderId == orderId {
			return c, nil
		}
	}
	return nil, errors.New("Order not found")
}
func (c *orderList) FindByPaidStatus(billStatus bool) ([]models.Order, error) {
	var listOrderByPaidStatus []models.Order
	for _, c := range c.orders {
		if c.IsPaid == billStatus {
			listOrderByPaidStatus = append(listOrderByPaidStatus, *c)
		}
	}
	if len(listOrderByPaidStatus) == 0 {
		return nil, errors.New("order is empty")
	}
	return listOrderByPaidStatus, nil
}

func NewOrderList() OrderList {
	return &orderList{}
}
