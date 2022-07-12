package main

import (
	"fmt"
)

const (
	a = 2020
	b = a + iota
	c = a + iota
	d = a + iota
)

func main() {
	fmt.Printf("%d Su tipo es %q\n", a, getTipo(a))
	fmt.Printf("%d Su tipo es %q\n", b, getTipo(b))
	fmt.Printf("%d Su tipo es %q\n", c, getTipo(c))
	fmt.Printf("%d Su tipo es %q\n", d, getTipo(d))
}

func getTipo(param1 interface{}) string {
	return fmt.Sprintf("%T", param1)
}