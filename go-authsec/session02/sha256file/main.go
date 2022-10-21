package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("sample-file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	h := sha256.New224()
	_, err = io.Copy(h, f)
	if err != nil {
		log.Fatalln("couldn't io.copy", err)
	}

	fmt.Printf("here's the type before Sum: %T\n", h)
	fmt.Printf("%x\n", h)
	xb := h.Sum(nil)
	fmt.Printf("here's the type after Sum: %T\n", xb)
	fmt.Printf("%x\n", xb)

	xb = h.Sum(nil)
	fmt.Printf("here's the type after second Sum: %T\n", xb)
	fmt.Printf("%x\n", xb)

	xb = h.Sum(nil)
	fmt.Printf("here's the type after third Sum and passing in xb: %T\n", xb)
	fmt.Printf("%x\n", xb)
}
