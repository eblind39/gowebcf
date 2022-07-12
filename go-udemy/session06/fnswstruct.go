package main

import (
	"fmt"
)

type persona struct {
	nombre string
	apellido string
	edad int
}

func (p persona) presentar() {
	fmt.Println("Hola, soy", p.nombre, p.apellido, "y tengo", p.edad)
}

func main() {
	p1 := persona{
		nombre: "Ernesto",
		apellido: "Gutierrez",
		edad: 27, 
	}
	p1.presentar()
}