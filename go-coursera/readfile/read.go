/*
 * Final Course Activity: read.go
 * Author: Ernesto Gutierrez
 * The goal of this assignment is to practice working with the ioutil and os packages in Go.
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// define a name struct which has two fields, fname and lname
type Employee struct {
	fname string `json:"fname" validate:"presence,min=1,max=20"`
	lname string `json:"lname" validate:"presence,min=1,max=20"`
}

func main() {
	var strFileName string
	var slcEmp []Employee

	// prompt the user for the name of the text file
	fmt.Print("Please enter the filename: ")
	strFileName = readString()

	file, err := os.Open(strFileName)
	if err != nil {
		log.Fatal("Can't read the file.")
		return
	}
	defer file.Close()

	// successively read each line of the text file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strLine := scanner.Text()
		arrData := strings.Split(strLine, " ")

		// create a struct which contains the first and last names found in the file
		empl := Employee{
			fname: arrData[0],
			lname: arrData[1],
		}
		// Each struct created will be added to a slice
		slcEmp = append(slcEmp, empl)
	}

	// iterate through your slice of structs and print the first and last names
	fmt.Println("The contents of the slice are: ")
	for i := 0; i < len(slcEmp); i++ {
		fmt.Println(slcEmp[i].fname, slcEmp[i].lname)
	}
}

// function to read any string for the input stream
func readString() string {
	var strTmp string

	// Use bufio to accept spaces into the string (i.e: "John Doe")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		strTmp = scanner.Text()
	}

	return strTmp
}
