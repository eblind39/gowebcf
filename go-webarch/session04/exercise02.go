package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	src := "kabir.txt" // set a different filename ir order to raise an error
	dst := "kabir-second.txt"
	err := copyFile(src, dst)
	var pe *os.PathError

	if errors.As(err, &pe) && errors.Is(err, os.ErrNotExist) {
		fmt.Printf("You need to provide the name %q of a valid file in your directory next to the executable - %s (%s)\n", src, pe.Err, pe.Path)
	} else if errors.As(err, &err) && errors.Is(err, os.ErrPermission) {
		fmt.Printf("You need permissions to open the name %q - %s\n", src, pe.Path)
	} else if err != nil {
		log.Panicln("In main, calling copyFile returned an error: ", err)
	}
}

func copyFile(src, dst string) error {
	fsrc, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Couldn't open file in copyFile: %w", err)
	}
	defer fsrc.Close()

	fdst, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("Couldn't create a file in copyFile: %w", err)
	}
	defer fdst.Close()

	n, err := io.Copy(fdst, fsrc)
	if err != nil {
		return fmt.Errorf("Couldn't copy a file in copyFile: %w", err)
	}

	fmt.Println("Just in development, nice to see bytes written: ", n)

	return nil
}
