/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 120, выход: 021
вход: 100, выход: 001
*/
package main

// Ваш код

import "fmt"

func main() {

	var number uint16
	fmt.Println("Введите трёхзначное число")
	fmt.Scan(&number)

	a := number / 100
	b := number / 10 % 10
	c := number % 10
	fmt.Println(c*100 + b*10 + a)
}
