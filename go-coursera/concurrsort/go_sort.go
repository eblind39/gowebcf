package main

import (
	"fmt"
	"runtime"
	"sync"
	"strings"
	"strconv"
)

func main() {
	var intSlice []int
	var wg sync.WaitGroup
	
	// read the numbers
	scanSliceElements(&intSlice)

	// print the original slice
	fmt.Println("\nThe original slice is: ")
	printSlice(intSlice)
	
	wg.Add(5)
	step := int(len(intSlice)/4)
	// order the slice using the BubbleSort algorithm
	go BubbleSort(intSlice[step*0:step*1], &wg)
	go BubbleSort(intSlice[step*1:step*2], &wg)
	go BubbleSort(intSlice[step*2:step*3], &wg)
	go BubbleSort(intSlice[step*3:], &wg)
	// order the entire slice
	go BubbleSort(intSlice, &wg)
	wg.Wait()
	

	fmt.Println("\nThe ordered slice is: ")
	printSlice(intSlice)
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

/*
 * Order the slice from least to greatest using the Bubble Sort algorithm
 */
func BubbleSort(intSlice []int, wg *sync.WaitGroup) {
	intLength := len(intSlice)
	
	defer wg.Done()

	for i := 0; i < intLength-1; i++ {
		for j := 0; j < intLength-i-1; j++ {
			if intSlice[j] > intSlice[j+1] {
				Swap(intSlice, j)
			}
			runtime.Gosched()
		}
	}
	
	printSlice(intSlice)
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