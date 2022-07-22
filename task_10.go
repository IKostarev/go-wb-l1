package main

import "fmt"

func sorted(temp float32) string {
	x := int(temp / 10)  //делю на 10 чтоб оставить один знак перед запятой

	if temp < 0 {
		x--
	}

	return fmt.Sprintf("границы от %.1f до %.1f", float32(x * 10), float32(x + 1) * 10) //нахожу границы разделов температуры
}

func getTemp(temps []float32) map[string][]float32 {
	tempMap := make(map[string][]float32) 

	for _, i := range temps { //записываю в мапу границу + значения 
		sort := sorted(i)
		tempMap[sort] = append(tempMap[sort], i)
	}

	return tempMap
}

func main() {
	startArray := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	res := getTemp(startArray)

	for n, m := range res {
		fmt.Println("Результат", n, ":", m)
	}
}