package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	numbers := make([]int, 0, 3)
	var input string

	for {
		fmt.Print("Type an integer (or X to exit): ")
		_, _ = fmt.Scan(&input)

		if input == "X" || input == "x" {
			break
		}
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Error: %s is  not a valid number", input)
		}
		numbers = append(numbers, number)
		sort.Ints(numbers)

		fmt.Printf("numbers: %v\n", numbers)
	}
}
