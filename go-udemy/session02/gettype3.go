package main

import (
	"fmt"
)

func main() {
	a := `Esta es una cadena de
		múltiples líneas
		"Lo ves?"`

	fmt.Printf("%s\n Su tipo es %q\n", a, getTipo(a))
}

func getTipo(param1 interface{}) string {
	return fmt.Sprintf("%T", param1)
}