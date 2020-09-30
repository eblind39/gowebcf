package main

import (
	"fmt"
	"time"
)

var value int

//The race function increase to print at execute of routine with parallelism inside the time of milliseconds
func race() {
	value++
	fmt.Println(value)
}

func main() {
	go race()
	go race()
	time.Sleep(time.Millisecond)
}
