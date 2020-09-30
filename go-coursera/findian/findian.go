/*
 * Module 2 Activity: findian.go
 * Author: Ernesto Gutierrez
 * The program should print “Found!” for the following example entered strings,
 *		“ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
 * The program should print “Not Found!” for the following strings,
 *		“ihhhhhn”, “ina”, “xian”.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var strEval string

	fmt.Println("Please enter a string")
	// Use bufio to accept spaces into the string (i.e: "I d skd a efju N")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		strEval = scanner.Text()
	}

	if strings.ToUpper(string(strEval[0])) == string("I") && // the string starts with [i|I]?
		strings.Contains(strings.ToUpper(strEval), "A") && // AND the string contains [a|A]?
		strings.ToUpper(string(strEval[len(strEval)-1])) == "N" { // AND the string ends with [n|N]?
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
