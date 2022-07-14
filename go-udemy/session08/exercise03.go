package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Nombre   string
	Apellido string
	Edad     int
	Dichos   []string
}

func main() {
	u1 := User{
		Nombre:   "Eduar",
		Apellido: "Tua",
		Edad:     32,
		Dichos: []string{
			"Cachicamo diciéndole a morrocoy conchudo",
			"La mona, aunque se vista de seda, mona se queda",
			"Poco a poco se anda lejos",
		},
	}

	u2 := User{
		Nombre:   "Carlos",
		Apellido: "Pérez",
		Edad:     27,
		Dichos: []string{
			"Camarón que se duerme se lo lleva la corriente",
			"A ponerse las alpargatas que lo que viene es joropo",
			"No gastes pólvora en zamuro",
		},
	}

	u3 := User{
		Nombre:   "Che",
		Apellido: "López",
		Edad:     54,
		Dichos: []string{
			"Ni lava ni presta la batea",
			"Hijo de gato, caza ratón",
			"Más vale pájaro en mano que cien volando",
		},
	}

	usuarios := []User{u1, u2, u3}

	fmt.Println(usuarios)

	err := json.NewEncoder(os.Stdout).Encode(usuarios)
	if err != nil {
		fmt.Println(err)
	}
}
