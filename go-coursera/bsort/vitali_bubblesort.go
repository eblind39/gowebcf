/*
Write a Bubble Sort program in Go. The program should
prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line,
in sorted order, from least to greatest. Use your favorite
search tool to find a description of how the bubble sort
algorithm works.

As part of this program, you should write a function called
BubbleSort() which takes a slice of integers as an argument
and returns nothing. The BubbleSort() function should modify
the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap
operation which swaps the position of two adjacent elements in
the slice. You should write a Swap() function which performs this
operation. Your Swap() function should take two arguments, a slice
of integers and an index value i which indicates a position in the
slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position
i+1.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Converts a string of comma separated numerical characters
to integers and appends to the slice
*/
func StringToIntList(input string, intList *[]int) {
	textNums := strings.Split(input, ",")
	for i, val := range textNums {
		if i >= 10 { // Limit the input to 10 integers as required
			break
		}
		d, err := strconv.Atoi(val)
		if err == nil {
			*intList = append(*intList, d)
		}
	}
}

func Swap(intList *[]int, i int) {
	if i >= len(*intList)-1 {
		fmt.Println("Warning: length ", len(*intList), " index ", i)
	} else {
		(*intList)[i], (*intList)[i+1] =
			(*intList)[i+1], (*intList)[i]
	}
}

func BubbleSort(intList *[]int) {
	for i, _ := range *intList {
		for j, _ := range (*intList)[0 : len(*intList)-i-1] {
			if (*intList)[j] > (*intList)[j+1] {
				Swap(intList, j)
			}
		}
	}
}

func main() {
	var input string
	integers := make([]int, 0, 10)
	/* Your program should prompt the user for the name
	of the text file */
	fmt.Println("Please enter up to 10 integers separated by comma (no spaces):")
	fmt.Scan(&input)
	StringToIntList(input, &integers) // Convert input to list of integers
	BubbleSort(&integers)             // Sort
	fmt.Println(integers)             // Print
}
