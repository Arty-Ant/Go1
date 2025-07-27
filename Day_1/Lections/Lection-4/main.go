package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstBoolean bool
	fmt.Println(firstBoolean)

	aBoolean, bBoolean := true, true
	fmt.Println("AND:", aBoolean && bBoolean)
	fmt.Println("OR:", aBoolean || bBoolean)
	fmt.Println("NOT:", !aBoolean)

	var value int
	fmt.Println("введите число: ")
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("Number", value, "is even.")
	} else {
		fmt.Println("Number", value, "is odd.")
	}

	var color string
	fmt.Scan(&color)

	if strings.Compare(color, "green") == 0 {
		fmt.Println("Color is green")
	} else if strings.Compare(color, "red") == 0 {
	} else {
		fmt.Println("Unknown color")
	}

	if num := 10; num%2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}
	if width := 100; width > 100 {
		fmt.Println("width > 100")
		return
	}
	fmt.Println("width <= 100")

}
