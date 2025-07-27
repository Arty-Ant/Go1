/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1 9 2
Выход: 1 2 9

Про sort мы пока ничего не знаем!
Это задача на условный оператор
*/
package main

// Ваш код

import "fmt"

func main() {

	var number1, number2, number3 uint16
	fmt.Println("Введите первое число: ")
	fmt.Scan(&number1)
	fmt.Println("Введите второе число: ")
	fmt.Scan(&number2)
	fmt.Println("Введите третье число: ")
	fmt.Scan(&number3)
	fmt.Println("Вот ваши числа в порядке возрастания: ")
	if number1 <= number2 && number1 >= number3 {
		fmt.Println(number3, number1, number2)
	} else if number2 <= number1 && number2 >= number3 {
		fmt.Println(number3, number2, number1)
	} else if number3 <= number1 && number3 >= number2 {
		fmt.Println(number2, number3, number1)
	} else if number3 <= number2 && number3 >= number1 {
		fmt.Println(number1, number3, number2)
	} else if number1 <= number3 && number1 >= number2 {
		fmt.Println(number2, number1, number3)
	} else {
		fmt.Println(number1, number2, number3)
	}

}
