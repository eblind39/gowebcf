package main

import "github.com/eblind39/gowebcf/go-webarch/session04/fileWriter"

func main() {
	f := fileWriter.NewWriteFile("file.txt")
	f.WriteString("Hello World!")
	f.WriteString("More text")
	f.Close()

	/*err := f.Err()
	var fErr *writeFileError
	if errors.As(err, &fErr) {
		if fErr.op == "Close" {
			fmt.Println("Error happened while closing file")
		}
	}*/

	err := f.Err()
	if err != nil {
		panic(err)
	}
}
