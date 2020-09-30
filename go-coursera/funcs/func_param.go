package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, -9, 4, 9, 6, 7}

	fmt.Println(sumSliceElems(nums...))
	fmt.Println(getMax(nums...))
	fmt.Println(getMin(nums...))

	fmt.Println("-----------DIV------------")

	fmt.Println(applyIt(sumSliceElems, nums...))
	fmt.Println(applyIt(getMax, nums...))
	fmt.Println(applyIt(getMin, nums...))
}

func applyIt(afunct func(...int) int, vals ...int) int {
	return afunct(vals...)
}

func sumSliceElems(vals ...int) int {
	var suma int

	for _, val := range vals {
		suma += val
	}

	return suma
}

func getMax(vals ...int) int {
	var tmp int
	bfirst := true

	for _, val := range vals {
		if bfirst {
			tmp = val
			bfirst = false
		}
		if tmp < val {
			tmp = val
		}
	}

	return tmp
}

func getMin(vals ...int) int {
	var tmp int
	bfirst := true

	for _, val := range vals {
		if bfirst {
			tmp = val
			bfirst = false
		}
		if tmp > val {
			tmp = val
		}
	}

	return tmp
}
