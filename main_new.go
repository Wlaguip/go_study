package main

import "fmt"

func main() {
	var servarCount int = 5
	cpuLoad := 80.5
	var serverName string = "Main_iMac"
	isOnline := true

	if cpuLoad > 80 {
		fmt.Println("ВНИМАНИЕ...")
	} else {
		fmt.Println("Нагрузка в норме")
	}

	fmt.Printf("Сервер: %s\n", serverName)
	fmt.Printf("Количество: %d, Нагрузка: %.1f%%\n", servarCount, cpuLoad)
	fmt.Printf("Статус онлайн: %t\n", isOnline)

	const ramTotal int = 16
	var ramUsed float64 = 4.2
	res := float64(ramTotal) - ramUsed

	fmt.Printf("Сумма: %.2f", res)

	fmt.Println("\n--- Запуск проверки систем ---")

	for i := 5; i >= 1; i-- {
		fmt.Printf("Проверка сервера %d... OK\n", i)

	}
	fmt.Println("Пуск")

}
