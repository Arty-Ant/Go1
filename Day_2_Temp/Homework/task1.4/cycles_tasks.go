/*
=======
Задачи:
=======

3.4 В цикле получать от пользователя оценки по четырём экзаменам.

	Вывести сумму набранных им баллов.
	Функцию fmt.Scan() в коде используем только один раз.
*/
package main

// Ваш код

import "fmt"

func main() {
	var overall int
	fmt.Println("Введите оценки за ваши 4 экзамена через пробел:")
	for count, grade := 0, 0; count < 4; count++ {
		fmt.Scan(&grade)
		overall += grade
	}
	fmt.Println("Сумма ваших оценок =", overall, "баллов")
}
