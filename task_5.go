///-------НЕ ВЫПОЛНЕНА----------

package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const N = 15  //minute

type Mux struct {
	sync.Mutex
}

func main() {
	var arr = [10]int{0, 1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	var m Mux

	ch := make(chan int)

	timer := time.AfterFunc(N * time.Minute, func() {
		log.Println("Программа выполялась - ", N * time.Minute, " минут")
	})

	defer timer.Stop()

	for {
		go func () {			
			for _, i := range arr {
				ch <- i
				fmt.Println("Пишу в канал: ", i)
			}
			close(ch)
		}()

		wg.Add(2)

		go func () {
			for i := range ch {
				m.Lock()
				fmt.Println("Читаю из канала: ", <-ch)
				i++
				m.Unlock()
			}
			wg.Done()
		}()

		go func () {
			for i := range ch {
				m.Lock()
				fmt.Println("Читаю из 2 канала: ", <-ch)
				i++
				m.Unlock()
			}
			wg.Done()
		}()

		wg.Wait()

		break
	}

}