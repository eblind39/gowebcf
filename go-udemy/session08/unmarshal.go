// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"Edad"`
}

func main() {
	var s string = `[{"Nombre":"Ernesto","Apellido":"Gutierrez","Edad":32},{"Nombre":"Alvion","Apellido":"Hunter","Edad":31}]`
	bs := []byte(s)
	var persons []Persona

	err := json.Unmarshal(bs, &persons)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(persons)

	for i, v := range persons {
		fmt.Println("\nPersona #", i+1)
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
	}
}
