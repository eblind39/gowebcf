package main

import (
	"fmt"
)

func main() {
	// delete using slicing
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	
	y := append(x[:3], x[6:]...)
	fmt.Println(x)
	fmt.Println(y)
}