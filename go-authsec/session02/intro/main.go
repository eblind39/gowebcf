package main

import "fmt"

func main() {
	var y []int
	x := append(y, []int{1, 2, 3}...)
	fmt.Println(x)
}
