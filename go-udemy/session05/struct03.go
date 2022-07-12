package main

import (
	"fmt"
)

type vehiculo struct {
	puertas int
	color   string
}

type camion struct {
	vehiculo
	cuatroRuedas bool
}

type sedan struct {
	vehiculo
	lujoso bool
}

func main() {
	c1 := camion{
		vehiculo: vehiculo{
			puertas: 2,
			color:   "Blue",
		},
		cuatroRuedas: true,
	}

	s1 := sedan{
		vehiculo: vehiculo{
			puertas: 4,
			color:   "Red",
		},
		lujoso: true,
	}

	fmt.Println(c1)
	fmt.Println(c1.puertas, c1.vehiculo.color, c1.cuatroRuedas)

	fmt.Println(s1)
	fmt.Println(s1.vehiculo.puertas, s1.color, s1.lujoso)
}
