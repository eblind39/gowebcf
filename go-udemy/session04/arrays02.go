package main

import (
	"fmt"
)

func main() {
	// another solution
	x := []int{42, 42, 43, 44, 45}
	for i, v := range x{
		fmt.Printf("x[%d] = %d, %T\n", i, v, v)
	}
}