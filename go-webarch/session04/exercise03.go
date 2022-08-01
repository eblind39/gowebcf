package main

import (
	"errors"
	"fmt"
)

func main() {
	fooErr := fooFn()
	barErr := errors.Unwrap(fooErr)
	mooErr := errors.Unwrap(barErr)
	catErr := errors.Unwrap(mooErr)
	baseErr := errors.Unwrap(catErr)
	fmt.Printf("fooErr\n\t%s\n\n", fooErr)
	fmt.Printf("barErr\n\t%s\n\n", barErr)
	fmt.Printf("mooErr\n\t%s\n\n", mooErr)
	fmt.Printf("catErr\n\t%s\n\n", catErr)
	fmt.Printf("baseErr\n\t%s\n\n", baseErr)
}

func fooFn() error {
	return fmt.Errorf("This error is from FOO: %w", barFn())
}

func barFn() error {
	return fmt.Errorf("This error is from BAR: %w", mooFn())
}

func mooFn() error {
	return fmt.Errorf("This error is from MOO: %w", catFn())
}

func catFn() error {
	return errors.New("This error is from CAT")
}
