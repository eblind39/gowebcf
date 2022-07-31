package main

import (
	"errors"
	"fmt"
)

type ErrFile struct {
	Filename string
	Base     error
}

var ErrorNotExist = fmt.Errorf("File does not exist")

func (e ErrFile) Error() string {
	return fmt.Sprintf("File %s: %v", e.Filename, e.Base)
}

/*func (e ErrFile) Is(other error) bool {
	return other == ErrorNotExist
}*/

func (e ErrFile) Unwrap() error {
	return e.Base
}

func openFiledef(filename string) (string, error) {
	return "", ErrorNotExist
}

func openFiledef2(filename string) (string, error) {
	return "", ErrFile{
		Filename: filename,
		Base:     ErrorNotExist,
	}
}

func main() {
	_, err := openFiledef("test.txt")
	if err != nil {
		wrappedErrr := fmt.Errorf("Unable to open file %v: %w", "test.txt", err)
		if errors.Is(wrappedErrr, ErrorNotExist) {
			fmt.Println("This is an ErrNotExist")
		}
		fmt.Println(wrappedErrr)
	}

	_, err = openFiledef2("test.txt")
	if err != nil {
		if errors.Is(err, ErrorNotExist) {
			fmt.Println("This is an ErrorNotExist")
		}
		fmt.Println(err)
	}
}
