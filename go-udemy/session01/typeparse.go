package main

import (
	"fmt"
)

type mitipo int
var y int

func main() {
	var x mitipo
	
	fmt.Println(x);
	fmt.Printf("El tipo es: %T\n", x)
	
	x = 42
	fmt.Println(x)
	
	y := int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
