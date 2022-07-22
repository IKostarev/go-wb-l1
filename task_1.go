package main

import "fmt"

//инициализация структуры Human
type Human struct {
	//список полей
	Name string
	Age  int
}

//инициализация структуры Action
type Action struct {
	Human   //в структуре Action получаем доступ к полям и методам структуры Human
	isCoder bool
}

//методы структуры Human
func (h *Human) PrinterHuman() string {
	return fmt.Sprint("Name - ", h.Name, ", ", "Age - ", h.Age)
}

func main() {
	var a Action

	a.Name = "Ilya"
	a.Age = 22
	a.isCoder = true

	fmt.Println(a.PrinterHuman())
}
