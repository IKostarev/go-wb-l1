package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	sync.Mutex
}

func (c *Counter) Iteration() {
	c.Lock()
	c.count++
	c.Unlock()
}

func doIter(num int) {
	var wg sync.WaitGroup
	var counts Counter

	wg.Add(3)

	for i := 0; i < 3; i++ { //в цикле запускаю 3 горутины и считаю конкурентно до введеного значения
		go func() {
			for i := 0; i < (num / 3); i++ {
				counts.Iteration()
			}
			wg.Done()
		}()
	}

	if num%10 > 0 { //если поделилось не нацело, то учитываю это
		wg.Add(1)

		go func() {
			for i := 0; i < (num % 10); i++ {
				counts.Iteration()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Сумма получилась - ", counts.count)
}

func main() {
	num := 999

	doIter(num)
}
