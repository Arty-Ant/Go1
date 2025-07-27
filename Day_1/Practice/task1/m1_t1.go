package main

import "fmt"

func main() {
	var (
		age  int
		name string
	)
	fmt.Scan(&age)
	fmt.Scan(&name)
	fmt.Println(age-1, name)
}
