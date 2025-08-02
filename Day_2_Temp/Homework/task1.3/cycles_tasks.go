/*
=======
Задачи:
=======

3.3 Пользователь вводит число. Вывести таблицу умножения на это число (от 1 до 10)
*/
package main

// Ваш код

import "fmt"

func main() {
	var num int
	fmt.Println("Введите число:")
	fmt.Scan(&num)
	fmt.Println("Таблица умножения на это число от 1 до 10:")
	for multip := 1; multip <= 10; multip++ {
		fmt.Println(multip, "*", num, " = ", multip*num)
	}
}
