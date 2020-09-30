package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	first := "i"
	middle := "a"
	last := "n"
	fmt.Println("Enter the string")
	fmt.Scanln(&str)
	if strings.HasPrefix(str, first) {
		if strings.HasSuffix(str, last) {
			if strings.Contains(str, middle) {
				fmt.Println("Found")
			}
		}
	} else {
		fmt.Println("Not Found")
	}
}
