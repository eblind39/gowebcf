package main

import (
	"fmt"
)

func main() {
	s1 := struct {
		nombre  string
		amigos  map[string]int
		bebidas []string
	}{
		nombre: "Ernesto",
		amigos: map[string]int{
			"Alvison": 222,
			"Edwin":   333,
			"Condor":  777,
		},
		bebidas: []string{
			"Agua",
			"Chicha",
		},
	}

	fmt.Println("\nDatos generales:")
	fmt.Println(s1)
	fmt.Println(s1.nombre, s1.amigos, s1.bebidas)

	fmt.Println("\nListado de amigos:")
	for k, v := range s1.amigos {
		fmt.Println(k, v)
	}

	fmt.Println("\nListado de bebidas:")
	for i, v := range s1.bebidas {
		fmt.Println(i, v)
	}
}
