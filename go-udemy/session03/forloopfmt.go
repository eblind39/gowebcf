package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d\t%x\t%#U\t%q\n", i, i, i, fmt.Sprintf("%q", i))
	}
}