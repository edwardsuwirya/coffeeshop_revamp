package repositories

import (
	"enigmacamp.com/coffeeshop/models"
	"errors"
)

type CoffeeMenu interface {
	Register(coffee models.Coffee)
	GetCoffeeProduct(productName string) (models.Coffee, error)
	ViewProduct() []models.Coffee
}
type coffeeMenu struct {
	coffees []models.Coffee
}

func (c *coffeeMenu) Register(coffee models.Coffee) {
	c.coffees = append(c.coffees, coffee)
}

func (c *coffeeMenu) GetCoffeeProduct(productName string) (models.Coffee, error) {
	for _, c := range c.coffees {
		if c.ProductInfo() == productName {
			return c, nil
		}
	}
	return nil, errors.New("Product not found")
}

func (c *coffeeMenu) ViewProduct() []models.Coffee {
	return c.coffees
}

func NewCoffeeMenu() CoffeeMenu {
	menu := new(coffeeMenu)
	menu.Register(models.NewCappucino("Kopi Gila", "robusta", "almond"))
	menu.Register(models.NewEspresso("Kopi Jamu", "arabica", 1))
	menu.Register(models.NewEspresso("Kopi Racun", "arabica", 2))
	return menu
}
