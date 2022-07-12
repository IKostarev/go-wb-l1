package main

import (
	"errors"
	"fmt"
)

func invertBit(number int, bitNum int)(int,error){
	bitMask:=1<<(bitNum-1) //маска для получения числа с битом равным единица на позиции которую будем менять

	if len(fmt.Sprintf("%b", number)) < len(fmt.Sprintf("%b", bitMask)){ //проверяю что в битовой форме есть необходимый разряд
		return number, errors.New(fmt.Sprintf("В двоичной записи числа %v(%b) нет %v разряда", number, number, bitNum))
	}

	return number ^ bitMask, nil
}

func main() {
	number := 300  //исходное число
	i := 2         //номер бита для замены

	fmt.Printf("До смены бита:\t\t %v = %b \n", number, number)
	fmt.Printf("Бит для инверсии: %v\n", i)
	
	res, err := invertBit(number, i)
	
	if err != nil {
		fmt.Println("Ошибка при получении числа после замены", err)
		return
	}

	fmt.Printf("После смены бита: \t %v = %b \n", res, res)
}