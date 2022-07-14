package main

import (
	"encoding/json"
	"fmt"
)

type Mexican struct {
	Nombre   string   `json:"Nombre"`
	Apellido string   `json:"Apellido"`
	Edad     int      `json:"Edad"`
	Dichos   []string `json:"Dichos"`
}

func main() {
	s := `[{"Nombre":"Eduar","Apellido":"Tua","Edad":32,"Dichos":["Cachicamo diciéndole a morrocoy conchudo","La mona, aunque se vista de seda, mona se queda","Poco a poco se anda lejos"]},{"Nombre":"Carlos","Apellido":"Pérez","Edad":27,"Dichos":["Camarón que se duerme se lo lleva la corriente","A ponerse las alpargatas que lo que viene es joropo","No gastes pólvora en zamuro"]},{"Nombre":"M","Apellido":"Hmmmm","Edad":54,"Dichos":["Ni lava ni presta la batea","Hijo de gato, caza ratón","Más vale pájaro en mano que cien volando"]}]`

	var mexicanos []Mexican

	err := json.Unmarshal([]byte(s), &mexicanos)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(s)
	fmt.Println(mexicanos)

	for i, mexican := range mexicanos {
		fmt.Println("\nMexicano #", i+1)
		fmt.Println("Nombre: ", mexican.Nombre)
		fmt.Println("Apellido: ", mexican.Apellido)
		fmt.Println("Edad: ", mexican.Edad)
		fmt.Println("\t\tDichos de ", mexican.Nombre)
		for j, d := range mexican.Dichos {
			fmt.Println("\t\t", j+1, " - ", d)
		}
	}
}
