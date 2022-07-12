package main

import (
	"fmt"
)

type persona struct {
	nombre string
	apellido string
	saboresHeladoFav []string
}

func main() {
	p1 := persona{
		nombre: "Martin",
		apellido: "Perez",
		saboresHeladoFav: []string{
			"fresa",
			"chocolate",
			"napolitano",
		},
	}

	fmt.Println(p1)
	
	fmt.Print("\nSabores favoritos de helados:")
	for i, helado := range p1.saboresHeladoFav {
		fmt.Printf("\nhelado[%d] = %s", i, helado)
	}
}
