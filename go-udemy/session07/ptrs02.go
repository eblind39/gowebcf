package main

import (
	"fmt"
)

type persona struct {
	nombre string
	apellido string
	edad int
}

func main() {
	p1 := persona{
		nombre: "Jim",
		apellido: "Jose",
		edad: 21,
	}
	
	fmt.Println(p1)
	fmt.Println(&p1)
	fmt.Printf("\n%T - %T\n", p1, &p1)
	
	cambiame(&p1)
	fmt.Println(p1)
	fmt.Println(&p1)
	fmt.Printf("\n%T - %T\n", p1, &p1)
}

func cambiame(p *persona) {
	p.nombre = "JMichael"
	// Otra forma de acceder y modificar el campo del struct
	// (*p).nombre = "Michael"
}