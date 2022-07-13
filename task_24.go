package main

import (
	"fmt"
	"math"
)

type Point struct { //структура Поинт с координатами
	X, Y float64
}

func NewPoint(x, y float64) Point { //конструктор
	pnt := Point{x, y}
	return pnt
}

func (p *Point) Setter(x, y float64) {
	p.X = x
	p.Y = y
}

func (p *Point) Getter() (x, y float64) {
	return p.X, p.Y
}

func Calc(p1, p2 Point) float64 {
	x1, y1 := p1.Getter()
	x2, y2 := p2.Getter()

	long := math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))

	return long
}

func main() {
	a := NewPoint(0, 0)
	b := NewPoint(10, 15)

	fmt.Printf("Координаты первой точки: %v\n", a)
	fmt.Printf("Координаты второй точки: %v \n", b)
	fmt.Println("Расстояние между ними: ", Calc(a, b))
}
