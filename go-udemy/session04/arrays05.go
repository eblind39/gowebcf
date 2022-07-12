package main

import (
	"fmt"
)

func main() {
	estados := make([]string, 50, 50)
	
	estados = []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", 
"Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", 
"Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", 
"Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey",
"New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon",
"Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", 
"Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"}
	
	fmt.Println(estados, len(estados), cap(estados))
	
	for i, estado := range estados {
		fmt.Printf("estados[%d] = %s, pos=%d\n", i, estado, i+1)
	}
}