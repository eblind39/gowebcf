package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 4, 6, 7}

	fmt.Println(sumSliceElems(nums...))
}

func sumSliceElems(vals ...int) int {
	var suma int

	for _, val := range vals {
		suma += val
	}

	return suma
}
