package models

type Order struct {
	OrderId  string
	Orders   []Coffee
	IsServed bool
	IsPaid   bool
}

func NewOrder(orderId string, orders []Coffee) Order {
	return Order{
		orderId, orders, false, false,
	}
}
