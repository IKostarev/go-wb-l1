package main

import (
	"fmt"
	"sync"
)

func main() {
	// sync.Map решает проблему cache contention в стандартной библиотеке 
	// для таких случаев, когда	ключи в map стабильны (не обновляются часто)
	// и происходит намного больше чтений, чем записей
	var testMap sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ { // 3 воркера которые делаю 10 записей
		wg.Add(1)

		go func(id int) {
			for i := 0; i < 10; i++ {
				testMap.LoadOrStore(i, i)  // метод, позволяющий установить значение по ключу в случае если оно еще не установлено
			}

			fmt.Println(id)
			wg.Done()			
		}(i)
	}

	wg.Wait()
}