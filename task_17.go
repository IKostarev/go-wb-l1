package main

import (
	"fmt"
	"sort"
)

func binarySearch(slice []int, findInt int) (result bool) {
	sort.Ints(slice) //бинарный поиск работает только в отсортированном слайсе!!!

	lowKey := 0               // первый индекс
	highKey := len(slice) - 1 // последний индекс

	if (slice[lowKey] > findInt) || (slice[highKey] < findInt) {
		fmt.Println("Искомое число не в диапазоне слайса")
	}

	for lowKey <= highKey { // уменьшаем список рекурсивно

		mid := (lowKey + highKey) / 2 // середина

		if slice[mid] == findInt {
			result = true // мы нашли значение
			return
		}
		if slice[mid] < findInt {
			// если поиск больше середины - мы берем только блок с большими значениями увеличивая lowKey
			lowKey = mid + 1
			continue
		}
		if slice[mid] > findInt {
			// если поиск меньше середины - мы берем блок с меньшими значениями уменьшая highKey
			highKey = mid - 1
		}
	}
	return
}

func main() {
	sl := []int{15, 18, 9, 1, 14, 100}

	findInt := 500

	fmt.Printf("Искомое число - %v\n", findInt)
	fmt.Printf("Результат поиска - %v\n", binarySearch(sl, findInt))
}
