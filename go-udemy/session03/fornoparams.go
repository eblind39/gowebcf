package main

import (
	"fmt"
)

func main() {
	anionac := 1980
	anioactual := 2020
	for {
		if (anioactual - anionac ) >= 0 {
			fmt.Printf("%d\n", anionac)
		} else {
			break
		}
		anionac++
	}
}