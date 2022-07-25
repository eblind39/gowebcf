package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fr, err := os.Open("file-01.txt")
	if err != nil {
		fmt.Println("No file found on current directory", err)
		return
	}
	defer fr.Close()

	fw, err := os.Create("file-02.txt")
	if err != nil {
		fmt.Println("No file was created", err)
		return
	}
	defer fw.Close()

	n, err := io.Copy(fw, fr)
	if err != nil {
		fmt.Println("No content copied", err)
		return
	}
	fmt.Println("bytes written: ", n)
}
