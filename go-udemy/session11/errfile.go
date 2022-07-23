package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fhndlr, err := os.Create("names.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer fhndlr.Close()

	r := strings.NewReader("James Bond")
	io.Copy(fhndlr, r)
}
