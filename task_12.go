package main

import "fmt"

func createdSet(str []string) {
	testMap := make(map[string]string)
	var res []string

	for _, i := range str { 
		testMap[i] = i
	}

	for j, _ := range testMap {  //записываю в итоговый массив данные без повторений
		res = append(res, j)
	}

	fmt.Println(res)

}

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}   //исходные массив данных

	createdSet(str)
}