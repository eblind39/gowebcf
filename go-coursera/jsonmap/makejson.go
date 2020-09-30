/*
 * Module 4 Activity: makejson.go
 * Author: Ernesto Gutierrez
 * The goal of this program is to practice working with RFCs in Go, using JSON as an example.
 */

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	name    string
	address string
}

func main() {
	// create a map to add the name and address
	personDet := make(map[string]string)

	// prompts the user to first enter a name, and then enter an address
	// and add them using the keys “name” and “address”
	fmt.Print("Please enter the name: ")
	personDet["name"] = readString()
	fmt.Print("Please enter the address: ")
	personDet["address"] = readString()

	// use Marshal() to create a JSON object from the map
	personDetJSON, err := json.Marshal(personDet)
	if err != nil {
		fmt.Println("Error Marshaling the map")
	} else {
		strPersonData := string(personDetJSON)
		fmt.Println("The resulting JSON is: ")
		fmt.Println(strPersonData) // print the JSON object
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
