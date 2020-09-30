/*
 * Module 3 Activity: slice.go
 * Author: Ernesto Gutierrez
 * Using slices.
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// create an empty integer slice of size (length) 3
	var intSlice = make([]int, 3)

	intSlice = scanSliceElements(intSlice) // read the numbers
}

func scanSliceElements(intSlice []int) []int {
	var strNumber string
	var intCont int

	// "infinite" loop to read the items of the slice
	for {
		fmt.Println("Please enter an integer or 'X' to quit: ")
		fmt.Scan(&strNumber)

		// if the user enters 'x'/'X' break/stop the reading
		if strings.ToUpper(strNumber) == "X" {
			break
		}

		// attemp to convert the input as an integer
		intTmp, err := strconv.Atoi(strNumber)
		if err != nil {
			fmt.Println("Enter a valid integer") // error on convertion, it's not an integer
		} else {
			// if we reach the slice's length of three, then adjust to allow store +1 of its capacity
			if intCont >= 3 {
				// we use an axiliary slice to increase the capacity and swapping the content
				intSliceTmp := make([]int, len(intSlice)+1, cap(intSlice)+1)
				copy(intSliceTmp, intSlice)
				intSlice = intSliceTmp
				// another direct solution is using the append method
				// intSlice = append(intSlice, intTmp)
			}

			intSlice[intCont] = intTmp
			intSlice = orderSlice(intSlice, intCont+1) // order the slice
			printSlice(intSlice, intCont+1)            // print the ordered slice
			intCont++
		}
	}

	return intSlice
}

/*
 * Order the slice using the bubble sort algorithm
 */
func orderSlice(intSlice []int, intLength int) []int {
	var intTmp int

	for i := 0; i < intLength-1; i++ {
		for j := 0; j < intLength-i-1; j++ {
			if intSlice[j] > intSlice[j+1] {
				intTmp = intSlice[j]
				intSlice[j] = intSlice[j+1]
				intSlice[j+1] = intTmp
			}
		}
	}
	return intSlice
}

// print the ordered slice
func printSlice(intSlice []int, intLength int) {
	fmt.Println("The ordered slice is: ")
	for i := 0; i < intLength; i++ {
		fmt.Println(intSlice[i])
	}
}
