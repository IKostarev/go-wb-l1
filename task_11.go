package main

import (
	"fmt"
	"sort"
)

func getArr(a, b []int) {

	testMap := make(map[int]int)
	var res []int

	sort.Ints(a)   //дано что они неупорядоченные, поэтому сортирую их
	sort.Ints(b)

	for _, i := range a {
		testMap[i] = testMap[i] - 1
	}

	for _, j := range b {
		testMap[j] = testMap[j] + 1
	}

	for n, m := range testMap { //прохожусь по готовому массиву и если есть совпадение, записываю в итоговый массив
		if m == 0 {
			res = append(res, n)
		}
	}

	sort.Ints(res)  //сортирую итоговый массив

	fmt.Println(res)
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 8, 7, 6}  //тестовые множества
	arr2 := []int{10, 8, 2, 6, 4}

	getArr(arr1, arr2)
}