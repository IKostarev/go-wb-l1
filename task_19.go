package main

import "fmt"

func coap(str []rune) {
	fmt.Printf("Исходная строка: %v\n", string(str))

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}

	fmt.Printf("Итоговая строка: %v\n", string(str))
}

func main() {
	str := []rune("привет нормально ничего")

	coap(str)
}
