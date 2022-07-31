package main

import (
	"errors"
	"fmt"
	"time"
)

type ErrFileNotFound struct {
	Filename string
	When     time.Time
}

func (e ErrFileNotFound) Error() string {
	return fmt.Sprintf("File %s was not found at %v", e.Filename, e.When)
}

func (e ErrFileNotFound) Is(other error) bool {
	_, ok := other.(ErrFileNotFound)
	return ok
}

var ErrNotExist = fmt.Errorf("File does not exist")
var ErrUserNotExist = errors.New("User does not exist")

func openFile(filename string) (string, error) {
	return "", ErrNotExist
}

func main() {
	err := ErrUserNotExist
	if err == ErrUserNotExist {
		fmt.Println("You need to register first!")
	} else {
		fmt.Println("Unknow error")
	}

	errF := ErrFileNotFound{
		Filename: "test.txt",
		When:     time.Now(),
	}
	fmt.Println(errF)
	fmt.Println(errors.Is(errF, ErrFileNotFound{}))

	_, errOF := openFile("test.txt")
	if errOF != nil {
		wrappedErr := fmt.Errorf("Unable to open file %v, %w", "test.txt", errOF)
		if errors.Is(wrappedErr, ErrNotExist) {
			fmt.Println("This is an ErrNotExist")
		}
		panic(wrappedErr)
	}
}
