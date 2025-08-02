/*
Задача №1.
Программа получает на вход последовательность из 5 целых чисел.

Вам нужно определить вид последовательности:
 - возрастающая
 - убывающая
 - случайная
 - постоянная

В качестве ответа следуют выдать прописными латинскими буквами тип последовательности:
1. ASCENDING (строго возрастающая)
2. WEAKLY ASCENDING (нестрого возрастающая, то есть неубывающая)
3. DESCENDING (строго убывающая)
4. WEAKLY DESCENDING (нестрого убывающая, то есть невозрастающая)
5. CONSTANT (постоянная)
7. RANDOM (случайная)

Примеры входных и выходных данных:
In: 11 9 4 2 -1  Out: DESCENDING
In: 3 8 8 11 12  Out: WEAKLY ASCENDING
In: 2 -1 7 21 1  Out: RANDOM
In: 5 5 5 5 5     Out: CONSTANT

Подсказка: используем метод строки strings.Split()
*/

package main

// Ваш код

import (
	"fmt"
)

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)
	fmt.Scan(&e)
	switch {
	case a == b && b == c && c == d && d == e:
		fmt.Println("CONSTANT")
	case a < b && b < c && c < d && d < e:
		fmt.Println("ASCENDING")
	case a <= b && b <= c && c <= d && d <= e:
		fmt.Println("WEAKLY ASCENDING")
	case a > b && b > c && c > d && d > e:
		fmt.Println("DESCENDING")
	case a >= b && b >= c && c >= d && d >= e:
		fmt.Println("WEAKLY DESCENDING")
	default:
		fmt.Println("RANDOM")
	}
}
