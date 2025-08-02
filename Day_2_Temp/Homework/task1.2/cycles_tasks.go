/*
=======
Задачи:
=======

3.1 Пользователь вводит числа a и b (b больше a).

	Вывести все целые числа, расположенные между ними.

3.2 Доработать предыдущую задачу так, чтобы выводились только числа,

	делящиеся на 5 без остатка.
*/
package main

// Ваш код

import "fmt"

func main() {
	var num1, num2, tempNum int
	fmt.Println("Введите первое число:")
	fmt.Scan(&num1)
	fmt.Println("Введите второе число:")
	fmt.Scan(&num2)
	if num1 > num2 {
		tempNum = num1
		num1 = num2
		num2 = tempNum
	}
	fmt.Println("Между меньшим и большим числом находится числа кратные 5-ти: ")
	for num1 < num2 {
		num1++
		if num1%5 == 0 {
			fmt.Print(num1, " ")
		}

	}
}
