package main

import "fmt"

// ФУНКЦИЯ 1: Двигатель программы. Она запускается САМА.
func main() {
	cities := []string{"Moscow", "Dubai", "New York", "Tokyo", "Krasnodar"}
	cities = append(cities, "Sochi")

	fmt.Println("=== ОТЧЕТ ПОГОДЫ ===")

	for _, city := range cities {
		// Вот тут магия: мы вызываем вторую функцию и получаем из неё данные
		t, d := getWeather(city)

		fmt.Printf("Город: %-10s | Темп: %d | Погода: %s\n", city, t, d)
	}
}

// ФУНКЦИЯ 2: Справочник. Она НЕ запускается сама, её вызывает main.
func getWeather(city string) (int, string) {
	if city == "Moscow" {
		return 2, "Облачно"
	} else if city == "Dubai" {
		return 35, "Солнечно"
	}
	// Если города нет в списке, возвращаем это:
	return 0, "Нет данных"
}
