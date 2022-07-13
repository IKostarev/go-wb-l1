package main

import (
	"fmt"
)

func QuickSort(arr []int, low, high int) []int {
	l := low
	h := high

	center := arr[(low+high)/2]   //определяю центр

	for l <= h {
		for arr[l] < center {     //итерирую пока low не достигнет центра
			l++
		}

		for arr[h] > center {     //итерирую пока high не достигнет центра
			h--
		}

		if l <= h {
			arr[h], arr[l] = arr[l], arr[h]   //затем меняем местами
			l++
			h--
		}
	}

	if h > low {
		QuickSort(arr, low, h)   //если после верхнего цикла условие выполняется,
	}							 //то вызываем функцию которая будет проходить от low до h

	if l < high {
		QuickSort(arr, l, high)	 //аналогично верхнему условию, только это для правой части
	}

	return arr
}

func main() {
	startArray := []int{100, 150, 15, 0, 10, 23}

	fmt.Println("Начальный массив: ", startArray)

	lastArray := QuickSort(startArray, 0, len(startArray)-1)

	fmt.Println("Конечный массив: ", lastArray)
}
