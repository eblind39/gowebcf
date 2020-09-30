/*
 * Module 1 Activity: Bubble Sort Program
 * Author: Ernesto Gutierrez
 * Ordering slices.
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// create an empty integer slice
	var intSlice []int

	// read the numbers
	scanSliceElements(&intSlice)

	// print the original slice
	fmt.Println("\nThe original slice is: ")
	printSlice(intSlice)

	// order the slice using the BubbleSort algorithm
	BubbleSort(intSlice)

	// print the ordered slice
	fmt.Println("\nThe ordered slice is: ")
	printSlice(intSlice)
}

/*
 * read the numbers
 */
func scanSliceElements(intSlice *[]int) {
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
			// if we reach the slice's length of ten, then break the "infinite" loop
			// intSlice[intCont] = intTmp
			*intSlice = append(*intSlice, intTmp)
			intCont++
			if intCont == 10 {
				break
			}
		}
	}
}

/*
 * Order the slice from least to greatest using the Bubble Sort algorithm
 */
func BubbleSort(intSlice []int) {
	// var intTmp int
	intLength := len(intSlice)

	for i := 0; i < intLength-1; i++ {
		for j := 0; j < intLength-i-1; j++ {
			if intSlice[j] > intSlice[j+1] {
				Swap(intSlice, j)
			}
		}
	}
}

/*
 * Swap two continuous elements from a slice
 */
func Swap(intSlice []int, intPos int) {
	intTmp := intSlice[intPos]
	intSlice[intPos] = intSlice[intPos+1]
	intSlice[intPos+1] = intTmp
}

/*
 * print the slice
 */
func printSlice(intSlice []int) {
	intLength := len(intSlice)
	var strTmp string

	for i := 0; i < intLength; i++ {
		strTmp += strconv.Itoa(intSlice[i]) + ", "
	}
	strTmp = strTmp
	fmt.Println("[" + strTmp[0:len(strTmp)-2] + "]")
}
