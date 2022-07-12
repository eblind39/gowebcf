package main

import (
	"fmt"
)

const (
	a = 42
	b string = "FILE"
)

func main() {
	fmt.Printf("a = %v - Tipo es %q\n", a, getTipo(a))
	fmt.Printf("b = %v - Tipo es %q\n", b, getTipo(b))
}

func getTipo(param1 interface{}) string {
	return fmt.Sprintf("%T", param1)
}