package main

import (
	"fmt"
)

func main() {
	var arrTmp [5]int
	
	arrTmp[0] = 5
	arrTmp[1] = 7
	arrTmp[2] = 3
	arrTmp[3] = 9
	arrTmp[4] = 12
	
	for i, v := range arrTmp {
		fmt.Printf("arrTmp[%d] = %d, %T\n", i, v, v)
	}
	
	// another solution
	x := [5]int{42, 42, 43, 44, 45}
	for i, v := range x{
		fmt.Printf("x[%d] = %d, %T\n", i, v, v)
	}
}