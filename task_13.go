package main

import "fmt"

func main() {
	a := 10
	b := 22
	fmt.Printf("Число a = %v и число b = %v\n", a, b)

	a, b = b, a
	fmt.Printf("Число a = %v и число b = %v\n", a, b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("Число a = %v и число b = %v\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Printf("Число a = %v и число b = %v", a, b)

}