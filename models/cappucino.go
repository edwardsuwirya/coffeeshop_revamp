package models

import (
	"fmt"
	"sync"
	"time"
)

type Cappucino struct {
	ProductName string
	Coffee      string
	Milk        string
}

func (c Cappucino) Brew(status chan string) {
	var wg sync.WaitGroup
	wg.Add(2)
	status <- fmt.Sprintf("Creating %s at %v", c.ProductName, time.Now())
	go func() {
		time.Sleep(2 * time.Second)
		info := fmt.Sprintf("Grind %s", c.Coffee)
		status <- info
		defer wg.Done()
	}()

	go func() {
		time.Sleep(3 * time.Second)
		info := fmt.Sprintf("Milk %s is warmed", c.Milk)
		status <- info
		wg.Done()
	}()

	wg.Wait()
	time.Sleep(3 * time.Second)
	info := fmt.Sprintf("Pouring and serving %s at %v", c.ProductName, time.Now())
	status <- info
}
func (c Cappucino) ProductInfo() string {
	return c.ProductName
}

func NewCappucino(productName string, coffeeType string, milkType string) Coffee {
	return Cappucino{
		Coffee:      coffeeType,
		Milk:        milkType,
		ProductName: productName,
	}
}
