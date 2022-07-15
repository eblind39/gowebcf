package main

import (
	"fmt"
	"runtime"
)

// go install myarch.go
func main() {
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("Arch: ", runtime.GOARCH)
}
