package main

import (
	"fmt" //для несколких библиотек
)

func main() {
	var result float64
	var d, e float64 = 10, 4
	var check float64 = 20 / 3

	a := 42
	b := 20
	c := a / b
	result = d / e

	fmt.Printf("type of c %T \n", c)
	fmt.Printf("type of result %T and value of result %v\n", result, result)
	fmt.Printf("type of check %T and value of check %v\n", check, check)
}
