package main

import (
	"fmt"
	"strconv"
	"strings"
	// "strconv"
)

func swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}
func bubbleSort(input []int) {
	n := 10
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if input[i-1] > input[i] {
				fmt.Println("Swapping")
				swap(input, i)
				swapped = true
			}
		}
	}
	fmt.Println(input)
}

func main() {
	var input string
	// var arr []string
	fmt.Println("Enter a sequence of up to 10 intgers")
	fmt.Scanln(&input)
	var inputSlice []string = strings.Split(input, "")
	var intSlice [10]int
	for i := range inputSlice {
		copy := inputSlice[i]
		strconv.Atoi(copy)
		new := strconv.Atoi(copy)
		intSlice[i] = new
	}
	for i := range input {
		inputSlice[i] = strconv.Atoi(input[i])
		fmt.Println(input[i])
	}
	fmt.Println(inputSlice)
	bubbleSort(intSlice)

}
