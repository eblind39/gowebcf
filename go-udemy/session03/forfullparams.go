package main

import (
	"fmt"
)

func main() {
	fmt.Println("Números divisibles entre 4 en el rango [10-100]")
	for i := 10; i<=100; i++ {
		if i % 4 == 0 {
			fmt.Printf("%d - resto (%d)\n", i, i%4)
		}
	}
}
