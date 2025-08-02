/*
=======
Задачи:
=======

3.6 Получить от пользователя натуральное число, посчитать сумму цифр в нём.

	Решить с помощью индексов, т.е. работаем с числом, как со строкой.
*/
package main

// Ваш код

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var sum int
	var stringNum string
	fmt.Println("Введите случайное натуральное число:")
	fmt.Scan(&stringNum)
	stringNumsArray := strings.Split(stringNum, "")
	for _, value := range stringNumsArray {
		intNum, err := strconv.Atoi(value)
		sum += intNum
		if err != nil {
			fmt.Println("Ошибка преобразования:", err)
			return
		}
	}
	fmt.Println("Сумма введёных чисел равна:", sum)
}
