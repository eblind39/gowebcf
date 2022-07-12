package main

import (
	"fmt"
)

func main() {
	if x := 42; x == 42 {
		fmt.Println("Inicializando y verificando valor de variable");
	} else if x==42 {
		fmt.Println("Nueva condición de if");
	} else {
		fmt.Println("Otro valor de i diferente de 42 y 43");
	}
}
