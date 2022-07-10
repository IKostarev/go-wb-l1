package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	stopChannel := make(chan int)

	wg.Add(1)

	go func (stop <- chan int) {
		for {
			select {
			case <- stop:
				fmt.Println("Остановка горутины с помощью дополнительного канала")
				wg.Done()
				return
			default:
				fmt.Println("Еще работаю...")
				time.Sleep(5 * time.Second)
			}
		}
	}(stopChannel)

	time.Sleep(1 * time.Second)
	close(stopChannel)
}