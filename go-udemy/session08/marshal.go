// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func main() {
	p1 := Persona{
		Nombre:   "Ernesto",
		Apellido: "Gutierrez",
		Edad:     32,
	}

	p2 := Persona{
		Nombre:   "Alvion",
		Apellido: "Hunter",
		Edad:     31,
	}

	persons := []Persona{p1, p2}

	bs, err := json.Marshal(persons)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
	fmt.Fprintln(os.Stdout, string(bs))
	io.WriteString(os.Stdout, string(bs))
}
