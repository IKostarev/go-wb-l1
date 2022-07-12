package main

import "fmt"

var justString string

func createHudeString(n int) string {
	str := ""

	for i := 0; i <= n; i++ {
		str += "word "
	}

	return str
}

func someFunc() {
	v := ""
	v = createHudeString(1 << 10)

	justString = v[:100] //дан срез байт, если символ занимает больше чем 1 байт, то будет не совсем корректный результат
	fmt.Printf("Данный justString - %s\n", justString)

	justMyString := string([]rune(v)[:100]) //срез символов
	fmt.Printf("Мой justMyString - %s\n", justMyString)
}

func main() {
	someFunc()
}
