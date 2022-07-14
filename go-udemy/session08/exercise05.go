package main

import (
	"fmt"
	"sort"
)

type UsuarioX struct {
	Nombre   string
	Apellido string
	Edad     int
	Dichos   []string
}

type ByAge []UsuarioX
func (a ByAge) Len() int { return len((a)) }
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

type ByLastName []UsuarioX
func (a ByLastName) Len() int { return len((a)) }
func (a ByLastName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByLastName) Less(i, j int) bool { return a[i].Apellido < a[j].Apellido }

func main() {
	u1 := UsuarioX{
		Nombre:   "Eduar",
		Apellido: "Tua",
		Edad:     32,
		Dichos: []string{
			"Cachicamo diciéndole a morrocoy conchudo",
			"La mona, aunque se vista de seda, mona se queda",
			"Poco a poco se anda lejos",
		},
	}

	u2 := UsuarioX{
		Nombre:   "Carlos",
		Apellido: "Pérez",
		Edad:     27,
		Dichos: []string{
			"Camarón que se duerme se lo lleva la corriente",
			"A ponerse las alpargatas que lo que viene es joropo",
			"No gastes pólvora en zamuro",
		},
	}

	u3 := UsuarioX{
		Nombre:   "Che",
		Apellido: "López",
		Edad:     54,
		Dichos: []string{
			"Ni lava ni presta la batea",
			"Hijo de gato, caza ratón",
			"Más vale pájaro en mano que cien volando",
		},
	}

	usuarios := []UsuarioX{u1, u2, u3}
	fmt.Println(usuarios)
	sort.Sort(ByAge(usuarios))
	fmt.Println("\t\tSorted by Age --->")
	fmt.Println(usuarios)

	sort.Sort(ByLastName(usuarios))
	fmt.Println("\t\tSorted by LastName --->")
	fmt.Println(usuarios)
}
