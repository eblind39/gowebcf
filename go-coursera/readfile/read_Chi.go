package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Name struct
type Name struct {
	fname string
	lname string
}

func main() {
	// reader := bufio.NewReader(os.Stdin)
	names := make([]Name, 0, 3)

	fmt.Println("What is the name of your text file?")
	fileName := readString()
	if filepath.Ext(fileName) != ".txt" {
		fileName = fileName + ".txt"
	} else {
		fileName = fileName
	}

	fileName = strings.Trim(fileName, "\n")

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		var name Name
		name.fname, name.lname = s[0], s[1]
		names = append(names, name)
	}

	for _, line := range names {
		fmt.Println(line)
	}
}

func readString() string {
	var strTmp string

	// Use bufio to accept spaces into the string (i.e: "John Doe")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		strTmp = scanner.Text()
	}

	return strTmp
}
