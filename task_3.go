package main

import (
	"fmt"
	"sync"
)

type Amont struct { //структура для сохранения суммы и инициализации Мьютекса
	int //поле для записи суммы
	sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	var s Amont //экземпляр структуры

	arr := [5]int{2, 4, 6, 8, 10} //исходный массив

	ch := make(chan int) //создание канала

	go func() {
		for _, i := range arr {
			pow := i * i //вычисляем квадраты
			ch <- pow    //записываем квадраты в канал
		}
		close(ch)
	}()

	wg.Add(2) //создаю 2 потока

	go func() { //первая функция для конкурентного вычисления квадрата
		for i := range ch {
			s.Lock()   //блокирую доступ к Мьютексу
			s.int += i //приплюсовываю результат
			s.Unlock() //разблокирую доступ к Мьютексу
			fmt.Println("Посчитал значение для - ", i)
		}
		wg.Done() //закрываю поток
	}()

	go func() { //вторая функция для конкурентного вычисления квадрата
		for i := range ch {
			s.Lock()
			s.int += i
			s.Unlock()
			fmt.Println("Посчитал значение для - ", i)
		}
		wg.Done() //закрываю поток
	}()

	wg.Wait() //ждем пока все открытые потоки закроются - синхронизируем потоки
	fmt.Println("Answer - ", s.int)

}
