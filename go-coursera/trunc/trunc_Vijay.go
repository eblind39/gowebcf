package main

import (
	"fmt"

)

/* func main() {
	var x, y float64
	fmt.Printf("Enter the float value")
	fmt.Scanln(&x)
	fmt.println(math.Round(x))
	fmt.Printf("Enter another value")
	fmt.Scanln(&y)
	fmt.Println(math.Round(y))
} */



func main() {
    untruncated := 10 / 3.0
    truncated := float64(int(untruncated * 100)) / 100
    fmt.Println(untruncated, truncated)
}
