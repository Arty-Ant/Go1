package main

import "fmt"

func main() {
	a := 1
	fmt.Scan(&a)
	a1 := a / 100
	a2 := (a - a1*100) / 10
	a3 := a % (a1*100 + a2*10)
	fmt.Println(a1, a2, a3)

	var number uint16
	fmt.Println("Введите трёхзначное число")
	fmt.Scan(&number)

	fmt.Println("Первая цифра: ", number/100)
	fmt.Println("Вторая цифра: ", number/10%10)
	fmt.Println("Третья цифра: ", number%10)
}
