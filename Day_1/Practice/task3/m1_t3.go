/*
Задача №3

На входе размер вклада(float64), кол-во лет(int64) и процент по вкладу(int64)
Проверить условия (от и до включительно):
вклад от 100 до 1_000_000
кол-во лет от 1 до 100
процент от 1 до 20

и посчитать размер вклада при выполнении условий.
В противном случае вывести сообщение о неправильных данных.

Использовать ежегодную капитализацию.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var vklad float64
	var years, perc int64
	fmt.Println("Введите размер вклада: ")
	fmt.Scan(&vklad)
	fmt.Println("Введите количество лет: ")
	fmt.Scan(&years)
	fmt.Println("Введите процент(%): ")
	fmt.Scan(&perc)
	if 100 <= vklad && 1000000 >= vklad && 1 <= years && 100 >= years && 1 <= perc && 20 >= perc {
		fmt.Println(vklad * (math.Pow((1 + (12*float64(years)/100)/12), float64(perc))))
		return
	}
	fmt.Println("Неправильные данные")
}
