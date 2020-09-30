/*
 * Module 2 Activity: trunc.go
 * Author: Ernesto Gutierrez
 * I tried to apply the various concepts we learned until now.
 */

package main

import (
	"fmt"
)

func main() {
	var fltNumber float32 // store the floating number entered
	var strNumber string  // store converted number to string

	fmt.Println("Please enter a float number")
	fmt.Scan(&fltNumber) // Scan the input stream

	// Convert the input to a string
	strNumber = fmt.Sprintf("%f", fltNumber)

	fmt.Println("The truncated number is: ")
	// call the customized function getIntPart to trunc the floating number
	// or extract the integer part
	fmt.Println(getIntPart(strNumber))

}

/*
 * Function to extract the integer part of the number
 */
func getIntPart(strFloat string) string {
	var strTmp string = "" // temporary string

	// iterate over the string which stores the floating number
	for i := 0; i < len(strFloat); i++ {
		var runChar = rune(strFloat[i]) // extract the character at position i as a rune
		if runChar == '.' {             // if the point (.) character was hit, -> then
			break // break the loop for
		}
		// on the other hand, concat the digits which belong to the integer part
		strTmp += string(runChar)
	}

	return strTmp // return the int part
}
