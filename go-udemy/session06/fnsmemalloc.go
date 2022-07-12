package main

import (
	"fmt"
)

func main() {
	a := foo() // mem allocate, x
	b := foo() // another mem allocate, x
	
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func foo() func() int {
	x := 1
	return func() int {
		x++
		return x
	}
}