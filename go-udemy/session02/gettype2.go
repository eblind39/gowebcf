package main

import (
	"fmt"
)

func main() {
	a := 42
	b := (a << 1)

	fmt.Printf("%d\t%b\t%x - Tipo es %q\n", a, a, a, getTipo(a))
	fmt.Printf("%d\t%b\t%x - Tipo es %q\n", b, b, b, getTipo(b))
}

func getTipo(param1 interface{}) string {
	return fmt.Sprintf("%T", param1)
}