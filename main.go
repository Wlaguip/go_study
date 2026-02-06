// main.go
package main

import "fmt"

func main_old() {
	stock := map[string]float64{
		"Молоко": 89.90,
		"Хлеб":   45.00,
		"Сыр":    210.50,
	}
	fmt.Println("--- Складской учет ---")

	fmt.Println("Цена на хлеб сегодня:", stock["Хлеб"], "руб.")

	product := "Колбаса"

	price, ok := stock[product]

	if ok {
		fmt.Printf("Товар %s найден. Цена: %.2f руб.\n", product, price)
	} else {
		fmt.Printf("Внимание! Товара %s нет в базе склада.\n", product)
	}

	total := 0.0
	for _, price := range stock {
		total = total + price
	}

	fmt.Printf("--- Итог ---\n")
	fmt.Printf("Общая стоимость товаров на складе: %.2f руб.\n", total)

	stock["Колбаса"] = 450.0
	delete(stock, "Хлеб")

	fmt.Println("\n--- Обновленный склад ---")
	for name, price := range stock {
		fmt.Printf("%s: %.2f руб.\n", name, price)
	}
}
