package main

import (
	"fmt"
)

type PersonY struct {
	Name string
}

func (p *PersonY) hablar() {
	fmt.Println("Hello, it's a pleasure to meet you! My name is: ", p.Name)
}

type Human interface {
	hablar()
}

func diAlgo(h Human) {
	h.hablar()
}

func main() {
	p := PersonY{Name: "Ernst"}
	diAlgo(&p)
}
