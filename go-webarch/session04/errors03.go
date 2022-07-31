package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("rumi_nf.txt")

	// Assertion
	var perr *os.PathError
	if errors.As(err, &perr) {
		fmt.Println("Here is info from patherror", perr.Op, perr.Path)
		errMsg := fmt.Errorf("Here is the original error %w and the path", err, perr.Path) // Returns the msg in a func
		fmt.Println(errMsg)
	}

	// better coding practice - choose this one
	if errors.Is(err, os.ErrPermission) {
		err = fmt.Errorf("You do not have permission to open the file: %w", err)
		log.Println(err)
	} else if errors.Is(err, os.ErrNotExist) {
		err = fmt.Errorf("The file does not exist: %w", err)
		log.Println(err)
	} else if err != nil {
		err = fmt.Errorf("File couldn't be opened: %w", err)
		log.Println(err)
	}

	// first approach
	if err == os.ErrPermission {
		err = fmt.Errorf("You do not have permission to open the file: %w", err)
		log.Println(err)
	} else if err == os.ErrNotExist {
		err = fmt.Errorf("The file does not exist: %w", err)
		log.Println(err)
	} else if err != nil {
		err = fmt.Errorf("File couldn't be opened: %w", err)
		log.Println(err)
	}

	defer f.Close()
}
