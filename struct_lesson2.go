package main

import "fmt"

type Animal struct {
	Species string
	Name    string
	Age     int
	Weight  float64
}

func main() {
	a1 := Animal{
		Species: "Корова",
		Name:    "Зорька",
		Age:     5,
		Weight:  450.5,
	}

	a2 := Animal{
		Species: "Бык",
		Name:    "Буян",
		Age:     3,
		Weight:  600.0,
	}

	fmt.Printf("Вид: %-10s | Кличка: %-8s | Возраст: %d лет | Вес: %.1f кг\n", a1.Species, a1.Name, a1.Age, a1.Weight)
	fmt.Printf("Вид: %-10s | Кличка: %-8s | Возраст: %d лет | Вес: %.1f кг\n", a2.Species, a2.Name, a2.Age, a2.Weight)

}
