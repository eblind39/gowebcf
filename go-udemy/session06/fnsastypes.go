package main

import (
	"fmt"
)

var x int
var g func()

func main() {

	f := func() {
		for i := 0; i <= 2; i++ {
			fmt.Println(i)
		}
	}
	
	f()
	fmt.Printf("%T\n", f)
	
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	
	g := f
	g()
	fmt.Printf("Esta es g %T\n", g)

	fmt.Println("Listo!")
}
