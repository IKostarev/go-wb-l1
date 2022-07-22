package main

import (
	"fmt"
	"sync"
)

func main() {
	var inputChan = make(chan int)
	var outputChan = make(chan int)
	var wg sync.WaitGroup

	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	wg.Add(3)

	go func(inputChan chan int) { //первый канал для записи чисел из массива
		for _, i := range x {
			inputChan <- i
			//fmt.Println("Пишу в первый канал - ", i)
		}
		close(inputChan)
		wg.Done()
	}(inputChan)

	go func(inputChan chan int, outputChan chan int) { //второй канал для записи из входного канала, преобразование числа и запись в выходной канал
		for i := range inputChan {
			outputChan <- i * 2
			//fmt.Println("Пишу во второй канал - ", i * 2)
		}
		close(outputChan)
		wg.Done()
	}(inputChan, outputChan)

	go func(outputChan chan int) { //читаю из выходного канала
		for i := range outputChan {
			fmt.Printf("Читаю из второго канала значение - %v", i)
			fmt.Printf("\n")
		}
		wg.Done()
	}(outputChan)

	wg.Wait()

}
