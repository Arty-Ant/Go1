/*
Задача №1
Вводим любое натуральное число.
Нужно посчитать сумму цифр числа, с помощью цикла for

Пример:
In: 4521
Out: 12
*Задание 1.1: 4 + 5 + 2 + 1 = 12 - добавить к выводу сумму как выражение
*/
package main

// Ваш код

import (
	"fmt"
)

func main() {
	var number, sum int
	var statement string

	fmt.Scan(&number)
	for number != 0 {
		digit := number % 10
		if len(statement) == 0 {
			statement = fmt.Sprintf("%d", digit) + " = "
		} else {
			statement = fmt.Sprintf("%d + ", digit) + statement
		}

		sum += digit //sum = 0
		number /= 10
	}
	fmt.Println("сумма числа:", sum)
	fmt.Printf("%s%v\n", statement, sum)
}
