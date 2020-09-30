package main

import (
	"fmt"
)

func main() {
	v := applyIt(func(x int) int {
		return 2 * x
	}, 5)
	fmt.Println(applyIt(incFn, 2))
	fmt.Println(applyIt(decFn, 2))
	fmt.Println(v)
}

func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func incFn(x int) int { return x + 1 }
func decFn(x int) int { return x - 1 }
