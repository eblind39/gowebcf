package main

import (
	"fmt"
)

func main() {
	personas := map[string][]string{
		`eduar_tua`: {`computadoras`, `montaña`, `playa`},
		`carlos_ramon`: {`leer`, `comprar`, `música`},
		`juan_bimba`: {`helado`, `pintar`, `bailar`},
	}
	
	// adding a new item to the map
	personas["Jim"] = []string{`Running`, `Hiking`, `Surfing`}
	
	for llave, valor := range personas  {
		fmt.Printf("personas[%s] = %s\n", llave, valor)
		for i, hobbie := range valor {
			fmt.Printf("\thobbie[%d] = %s\n", i, hobbie)
		}
	}
}