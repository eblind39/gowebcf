package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// Almacenando el valor de impresión en s
	s := fmt.Sprintf("%d\t%s\t%t", x, y, z)

	fmt.Println(s)
}