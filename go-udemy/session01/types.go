package main

import (
	"fmt"
)

type mitipo int

func main() {
	var x mitipo
	
	fmt.Println(x);
	fmt.Printf("El tipo es: %T\n", x)
	
	x = 42
	fmt.Println(x);
}
