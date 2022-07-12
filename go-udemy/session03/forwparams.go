package main

import (
	"fmt"
)

func main() {
	anionac := 1980
	anioactual := 2020
	for (anioactual - anionac ) >= 0 {
		fmt.Printf("%d\n", anionac)
		anionac++
	}
}