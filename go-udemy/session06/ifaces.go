package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type cuadrado struct {
	lado float64
}

type circulo struct {
	radio float64
}

func (c cuadrado) area() float64 {
	return math.Pow(c.lado, 2)
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.radio, 2)
}

func info(f forma) {
	fmt.Println("Mi área es", f.area())
}

func main() {
	sqr := cuadrado{
		lado: 2.3,
	}
	crcl := circulo{
		radio: 4.5,
	}

	info(sqr)
	info(crcl)
}
