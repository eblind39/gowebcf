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

	xs := []shape{s1, s2}

	bs, err := json.Marshal(xs)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(bs))
}
