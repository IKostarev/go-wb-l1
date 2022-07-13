package main

import (
	"fmt"
	"strings"
)

func unique(str string) bool {
	str = strings.ToLower(str) //привожу все к нижнему регистру

	r := []rune(str) //руна чтобы учесть юникод

	mapUn := make(map[rune]int)

	for _, i := range r { //если одинаковые символы, плюсанет два и больше раз и через условие нахожу это
		mapUn[i]++

		if mapUn[i] > 1 {
			return true
		}
	}

	return false
}

func main() {
	str := "aabcd"

	fmt.Println("Начальная строка - ", str)

	fmt.Println("Есть повтор символов? - ", unique(str))
}
