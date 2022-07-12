package main

import (
	"fmt"
	"reflect"
)

func definVarType(getInterface interface{}) string {  //на входе получаем интерфейс с переменной и дальше через встроенную функцию reflect определяем тип
	return fmt.Sprintln("Тип переменной: ", reflect.TypeOf(getInterface).String())
}

func main() {
	num := 1
	flo := 3.1
	str := "word"
	ch := make(chan int)
	bol := true

	fmt.Printf("Переменная %v ", num)
	fmt.Println(definVarType(num))
	fmt.Printf("\n")

	fmt.Printf("Переменная %v ", flo)
	fmt.Println(definVarType(flo))
	fmt.Printf("\n")

	fmt.Printf("Переменная %v ", str)
	fmt.Println(definVarType(str))
	fmt.Printf("\n")

	fmt.Printf("Переменная %v ", ch)
	fmt.Println(definVarType(ch))
	fmt.Printf("\n")

	fmt.Printf("Переменная %v ", bol)
	fmt.Println(definVarType(bol))
	fmt.Printf("\n")
}