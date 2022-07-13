package main

import (
	"fmt"
	"strings"
)

func coap(str string) string {
	newStr := strings.Split(str, " ") //разделяю строку по пробелам

	for i := 0; i < (len(newStr) / 2); i++ {
		newStr[i], newStr[len(newStr)-i-1] = newStr[len(newStr)-i-1], newStr[i]
	}

	return strings.Join(newStr, " ") //соединяем обратно
}

func main() {
	str := "hello every one"

	fmt.Printf("Исходная строка: %v\n", str)
	fmt.Printf("Итоговая строка: %v\n", coap(str))
}
