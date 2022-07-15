package main

import (
	"fmt"
	"math"
)

type circulo struct {
	radio float64
}

func (c *circulo) area() float64 {
	return math.Pi * math.Pow(c.radio, 2)
}

type forma interface {
	area() float64
}

func info(s forma) {
	fmt.Println("area", s.area())
}

func main() {
	c := circulo{radio: 5}
	info(&c)
}
