package delivery

import (
	"enigmacamp.com/coffeeshop/models"
	"enigmacamp.com/coffeeshop/repositories"
	"enigmacamp.com/coffeeshop/usecases"
	"fmt"
)

type CoffeeShopApp struct {
	takeOrder         usecases.TakeOrder
	kitchenProduction usecases.Production
	findMenu          usecases.FindMenu
	billing           usecases.Bill
}

func (csp CoffeeShopApp) TakeOrder(coffeeOrder []models.Coffee) {
	order := csp.takeOrder.CreateCustomerOrder(coffeeOrder)
	csp.sendToKitchen(order)
}
func (csp CoffeeShopApp) sendToKitchen(order *models.Order) {
	csp.kitchenProduction.Produce(order)
}
func (csp CoffeeShopApp) ChooseCoffee(productName string) models.Coffee {
	product, err := csp.findMenu.FindByName(productName)
	if err != nil {
		return nil
	}
	return product
}

func (csp CoffeeShopApp) ViewUnpaidBill() {
	fmt.Println(csp.billing.GetListUnpaid())
}
func (csp CoffeeShopApp) PrintMenu() {
	fmt.Println(csp.findMenu.ViewAll())
}

func (csp CoffeeShopApp) PrintBill(orderId string) {
	err := csp.billing.CreateBill(orderId)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func NewCoffeeShopApp() CoffeeShopApp {
	orderList := repositories.NewOrderList()
	return CoffeeShopApp{
		takeOrder:         usecases.NewTakeOrder(orderList),
		kitchenProduction: usecases.NewKitchenProduction(),
		findMenu:          usecases.NewFindMenu(repositories.NewCoffeeMenu()),
		billing:           usecases.NewBill(orderList),
	}
}
