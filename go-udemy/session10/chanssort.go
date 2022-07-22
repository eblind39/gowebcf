package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var intSlice []int
	
	// read the numbers
	scanSliceElements(&intSlice)

	// print the original slice
	fmt.Println("\nThe original slice is: ")
	printSlice(intSlice)
	fmt.Println("")
	
	// Use channels to send and receive the ordered slice
	step := int(len(intSlice)/4)
	intSlice14CH := make(chan []int)
	intSlice24CH := make(chan []int)
	intSlice34CH := make(chan []int)
	intSlice44CH := make(chan []int)
	// order the slice using the BubbleSort algorithm
	go BubbleSort(intSlice[step*0:step*1], intSlice14CH)
	go BubbleSort(intSlice[step*1:step*2], intSlice24CH)
	go BubbleSort(intSlice[step*2:step*3], intSlice34CH)
	go BubbleSort(intSlice[step*3:], intSlice44CH)
	intSlice14 := <- intSlice14CH
	intSlice24 := <- intSlice24CH
	intSlice34 := <- intSlice34CH
	intSlice44 := <- intSlice44CH
	fmt.Println("Ordered 1/4 is ", intSlice14)
	fmt.Println("Ordered 2/4 is ", intSlice24)
	fmt.Println("Ordered 3/4 is ", intSlice34)
	fmt.Println("Ordered 4/4 is ", intSlice44)
	
	// merge 4 sorted slices
	var intSliceOS []int
	intSliceOS = append(intSliceOS, intSlice14...)
	intSliceOS = append(intSliceOS, intSlice24...)
	intSliceOS = append(intSliceOS, intSlice34...)
	intSliceOS = append(intSliceOS, intSlice44...)
	
	// order the entire slice
	intSliceFULLCH := make(chan []int)
	go BubbleSort(intSliceOS, intSliceFULLCH)
	intSliceFULL := <- intSliceFULLCH

	fmt.Println("\nThe ordered slice is: ")
	printSlice(intSliceFULL)
}

/*
 * Order the slice from least to greatest using the Bubble Sort algorithm
 */
func BubbleSort(intSlice []int, intSliceRet chan []int) {
	intLength := len(intSlice)

	for i := 0; i < intLength-1; i++ {
		for j := 0; j < intLength-i-1; j++ {
			if intSlice[j] > intSlice[j+1] {
				Swap(intSlice, j)
			}
		}
	}
	
	intSliceRet <- intSlice
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

/*
 * read the numbers
 */
func scanSliceElements(intSlice *[]int) {
	var strNumber string

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
			*intSlice = append(*intSlice, intTmp)
		}
	}
}