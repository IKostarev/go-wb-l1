package main

import (
	"fmt"
	"log"
)

func deleteSliceNumber(sl []int, n int) []int {
	if n > len(sl) {
		log.Fatalln("Ввели неправильный элемент для удаления") //логирую ошибку
	}

	sl = append(sl[:n-1], sl[n:]...)

	return sl
}

func main() {
	sl := []int{15, 20, 25, 30, 35} //начальный слайс
	n := 33                         //элемент для удаления

	fmt.Println("Начальный слайс: ", sl)
	fmt.Println("Удаляем элемент под номером: ", n)
	fmt.Println("Результат: ", deleteSliceNumber(sl, n))
}
