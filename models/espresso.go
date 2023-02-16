package models

import (
	"fmt"
	"sync"
	"time"
)

type Espresso struct {
	ProductName string
	Coffee      string
	Shot        int
}

func (e Espresso) Brew(status chan string) {
	var wg sync.WaitGroup
	wg.Add(1)
	status <- fmt.Sprintf("Creating %s at %v", e.ProductName, time.Now())
	go func() {
		time.Sleep(time.Duration(2*e.Shot) * time.Second)
		info := fmt.Sprintf("Grind %s", e.Coffee)
		status <- info
		defer wg.Done()
	}()
	wg.Wait()
	time.Sleep(3 * time.Second)
	info := fmt.Sprintf("Pouring and serving %s at %v", e.ProductName, time.Now())
	status <- info
}
func (e Espresso) ProductInfo() string {
	return e.ProductName
}

func NewEspresso(productName string, coffeeType string, shot int) Coffee {
	return Espresso{
		Coffee:      coffeeType,
		Shot:        shot,
		ProductName: productName,
	}
}
