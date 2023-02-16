package main

import (
	"enigmacamp.com/coffeeshop/delivery"
	"enigmacamp.com/coffeeshop/models"
)

func main() {
	app := delivery.NewCoffeeShopApp()
	app.PrintMenu()
	coffee1 := app.ChooseCoffee("Kopi Gila")
	coffee2 := app.ChooseCoffee("Kopi Jamu")
	app.TakeOrder([]models.Coffee{coffee1, coffee2})
	app.ViewUnpaidBill()
	app.PrintBill("123")
	app.ViewUnpaidBill()
}
