package usecases

import (
	"enigmacamp.com/coffeeshop/models"
	"enigmacamp.com/coffeeshop/repositories"
)

type FindMenu interface {
	FindByName(productName string) (models.Coffee, error)
	ViewAll() []models.Coffee
}

type findMenu struct {
	coffeeMenu repositories.CoffeeMenu
}

func (fm findMenu) FindByName(productName string) (models.Coffee, error) {
	return fm.coffeeMenu.GetCoffeeProduct(productName)
}
func (fm findMenu) ViewAll() []models.Coffee {
	return fm.coffeeMenu.ViewProduct()
}

func NewFindMenu(coffeeMenu repositories.CoffeeMenu) FindMenu {
	return findMenu{coffeeMenu}
}
