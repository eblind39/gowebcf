package main

import (
	"fmt"
)
func main() {
	var x float64
	var y int32
	fmt.Printf("Please enter a floating point number: ")
	fmt.Scan(&x)
	y = int32(x)
	fmt.Println("The truncated integer is:", y)
}