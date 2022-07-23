package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type persona struct {
	Nombre   string
	Apellido string
	Frases   []string
}

func main() {
	p1 := persona{
		Nombre:   "James",
		Apellido: "Bond",
		Frases:   []string{"Shaken, not stirred", "¿Algún último deseo?", "Nunca digas nunca."},
	}

	bs, err := aJSON(p1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))

	bs, err = aJSON2(p1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}

// aJSON también necesita retorna un error
func aJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("An error happened in aJSON: %v", err)
	}
	return bs, nil
}

func aJSON2(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, errors.New(fmt.Sprintf("An error happened in aJSON: %v", err))
	}
	return bs, nil
}
