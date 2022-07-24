package main

import (
	"fmt"
)

type human interface {
	speak()
}

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("I'm a person - this is my name: ", p.first)
}

type scientist struct {
	person
	ltk bool
}

func (sa scientist) speak() {
	fmt.Println("I'm a scientist - this is my name: ", sa.first)
}

func main() {
	p1 := person{
		first: "Karl Jung",
	}

	sc1 := scientist{
		person: person{first: "Blaise Pascal"},
		ltk:    true,
	}

	var x, y human
	x = p1
	y = sc1
	x.speak()
	y.speak()

	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", sc1)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}
