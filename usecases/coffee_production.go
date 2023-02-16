package usecases

import (
	"enigmacamp.com/coffeeshop/models"
	"fmt"
	"sync"
)

type Production struct {
}

func (p Production) Produce(orders *models.Order) {
	status := make(chan string)

	var orderWg sync.WaitGroup
	for _, o := range (*orders).Orders {
		orderWg.Add(1)
		go func(c models.Coffee) {
			c.Brew(status)
			defer orderWg.Done()
		}(o)
	}

	go func() {
		orderWg.Wait()
		close(status)
	}()

	for s := range status {
		fmt.Println(s)
	}
	orders.IsServed = true
}

func NewKitchenProduction() Production {
	return Production{}
}
