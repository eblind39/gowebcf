package main

import (
	"errors"
	"fmt"
)

func main() {
	fooErr := foo()
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

func foo() error {
	return fmt.Errorf("This error is from FOO: %w", bar())
}

func bar() error {
	return fmt.Errorf("This error is from BAR: %w", moo())
}

func moo() error {
	return fmt.Errorf("This error is from MOO: %w", cat())
}

func cat() error {
	return errors.New("This error is from CAT")
}
