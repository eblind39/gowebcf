package main

import (
	"fmt"
	"io"
	"os"
)

// --------Error helper
type writeFileError struct {
	op  string
	err error
}

func (w writeFileError) Error() string {
	return w.err.Error()
}

func (w writeFileError) Unwrap() error {
	return w.err
}

// --------File helper
type writeFile struct {
	f   *os.File
	err error
}

func newWriteFile(filename string) writeFile {
	f, err := os.Create(filename)
	return writeFile{
		f:   f,
		err: err,
	}
}

func (w *writeFile) writeString(text string) {
	if w.err != nil {
		return
	}

	_, err := io.WriteString(w.f, text)
	if err != nil {
		w.err = writeFileError{
			op:  "writeString",
			err: fmt.Errorf("Failed while writing a string: %w", err),
		}
	}
}

func (w *writeFile) Close() {
	if w.err != nil {
		return
	}

	err := w.f.Close()
	if err != nil {
		w.err = err
	}
}

// All errors returning from Err should be of type *writeFileError
func (w *writeFile) Err() error {
	return w.err
}

func main() {
	f := newWriteFile("file.txt")
	f.writeString("Hello World!")
	f.writeString("More text")
	f.Close()

	err := f.Err()
	if err != nil {
		panic(err)
	}
}
