package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(2 << 25)
	b := big.NewInt(2 << 25 + 10)
	fmt.Printf("Число a: %v\t b: %v\n", a, b)

	res := big.NewInt(0)

	res.Add(a, b)
	fmt.Printf("%v + %v = %v\n", a, b, res)

	res.Sub(a, b)
	fmt.Printf("%v - %v = %v\n", a, b, res)

	res.Mul(a, b)
	fmt.Printf("%v * %v = %v\n", a, b, res)

	res.Quo(a, b)
	fmt.Printf("%v / %v = %v\n", a, b, res)
}
