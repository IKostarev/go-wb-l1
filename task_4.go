///-------НЕ ВЫПОЛНЕНА----------

package main

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
)

func getKeyPressed() (ch rune, err error) {
	var (
        chChan  = make(chan rune, 1)
        errChan = make(chan error, 1)

        timer = time.NewTimer(5 * time.Minute)
    )
    defer timer.Stop()

	if err = keyboard.Open(); err != nil {
		return
	}

	defer keyboard.Close()

    go func(chChan chan<- rune, errChan chan<- error) {
        ch, _, err := keyboard.GetSingleKey()
        
		if err != nil {
            errChan <- err
            return
        }
		fmt.Println("Пишу в канал - ", ch)
        chChan <- ch
    }(chChan, errChan)

	fmt.Println("В канале chChan - ", chChan, "\nв канале errChan - ", errChan)

	select {
    case <-timer.C:
        return 0, err
    case ch = <-chChan:
    case err = <-errChan:
    }

    return
}

func main() {
	getKeyPressed()
}