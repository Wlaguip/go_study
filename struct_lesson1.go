package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func lesson1() {
	p1 := Product{
		Name:     "Молоко",
		Price:    89.50,
		Quantity: 10,
	}

	fmt.Printf("На складе: %s\n", p1.Name)
	fmt.Printf("Цена: %.2f руб. | В наличии: %d шт.\n", p1.Price, p1.Quantity)

	totalValue := p1.Price * float64(p1.Quantity)
	fmt.Printf("Общая стоимость партии: %.2f руб.\n", totalValue)

}
