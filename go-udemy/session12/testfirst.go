package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("La suma de 2 + 4 es: ", miSuma(2, 4))
	fmt.Println("La suma de 1 + 5 es: ", miSuma(1, 5))
	fmt.Println("La suma de 6 + 7 es: ", miSuma(6, 7))
}

func miSuma(xi ...int) int {
	var sum int

	for _, v := range xi {
		sum += v
	}

	return sum
	// return sum + 1 //Fail the test
}

func Cat(xs []string) string {
	s := xs[0]
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

func Join(xs []string) string {
	return strings.Join(xs, " ")
}
