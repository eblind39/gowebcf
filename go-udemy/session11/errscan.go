package main

import (
	"fmt"
)

func main() {
	var respuesta1, respuesta2, respuesta3 string

	fmt.Println("Name: ")
	_, err := fmt.Scan(&respuesta1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Favorite meal: ")
	_, err = fmt.Scan(&respuesta2)
	if err != nil {
		panic(err)
	}

	fmt.Println("Favorite sport: ")
	_, err = fmt.Scan(&respuesta3)
	if err != nil {
		panic(err)
	}

	fmt.Println(respuesta1, respuesta2, respuesta3)
}
