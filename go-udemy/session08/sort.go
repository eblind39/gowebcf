package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{5,2,6,3,1,4}	// unsorted
	strs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	sort.Ints(ints)
	fmt.Println(ints)

	sort.Strings(strs)
	fmt.Println(strs)

	for i, v := range strs {
		fmt.Println("\nString #", i+1)
		fmt.Println(v)
	}
}