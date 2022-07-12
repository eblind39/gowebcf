package main

import (
	"fmt"
)

func main() {
	p1 := []string{"Batman", "Jefe", "Vestido de negro"}
	p2 := []string{"Robin", "Ayudante", "Vestido de colores"}
	personas:= [][]string{p1, p2}
	
	fmt.Println(personas, len(personas), cap(personas))
	
	for i, persona:= range personas{
		fmt.Printf("personas[%d] = %s, pos=%d\n", i, persona, i+1)
		for j, data := range persona {
			fmt.Printf("\tpersona[%d] = %s, pos=%d\n", j, data , j+1)
		}
	}
}