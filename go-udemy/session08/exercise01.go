package main

import (
	"encoding/json"
	"fmt"
)

type Usuario struct {
	Nombre string
	Edad   int
}

func main() {
	var u1 Usuario = Usuario{
		Nombre: "Eduar",
		Edad:   32,
	}

	var u2 Usuario = Usuario{
		Nombre: "Juan",
		Edad:   27,
	}

	var u3 Usuario = Usuario{
		Nombre: "Alvison",
		Edad:   37,
	}

	var usuarios []Usuario = []Usuario{u1, u2, u3}

	bc, err := json.Marshal(usuarios)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bc))
}
