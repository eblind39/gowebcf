package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type shape struct {
	Name string
}

func main() {
	s1 := shape{
		Name: "triangle",
	}

	s2 := shape{
		Name: "square",
	}

	// -------Marshall
	xs := []shape{s1, s2}

	bs, err := json.Marshal(xs)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("JSON string: ", string(bs))

	// -------Unmarshall
	var xs2 []shape

	err = json.Unmarshal(bs, &xs2)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Back into a Go data structure", xs2)
}
