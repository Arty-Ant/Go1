/*
Задача №1
Вход:

	расстояние(50 - 10000 км),
	расход в литрах (5-25 литров) на 100 км и
	стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*Проверка условий по желанию
*/
package main

// Ваш код

import "fmt"

func main() {
	var dist, petrol float32
	const cost = 48
	fmt.Println("Расстояние в км.: ")
	fmt.Scan(&dist)
	fmt.Println("Расход в литрах: ")
	fmt.Scan(&petrol)
	if 50 <= dist && 10000 >= dist && 5 <= petrol && 25 >= petrol {
		fmt.Println(petrol / 100 * cost * dist)
		return
	}
	fmt.Println("Неправильные данные")
}
