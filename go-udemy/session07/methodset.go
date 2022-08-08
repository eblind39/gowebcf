package main

import (
	"fmt"
)

type T struct {
	name string
}

// Pointer type receiver
func (receiver *T) pointerMethod() {
	fmt.Printf("Pointer method called on \t%#v with address %p\n\n", *receiver, receiver)
}

// Value type receiver
func (receiver T) valueMethod() {
	fmt.Printf("Value method called on \t%#v with address %p\n\n", receiver, &receiver)
}

func main() {
	var (
		val     T  = T{}
		pointer *T = &val
	)

	fmt.Printf("Value created \t%#v with address %p\n", val, &val)
	fmt.Printf("Pointer created on \t%#v with address %p\n", *pointer, pointer)

	val.valueMethod()       // Implicitly converted to: (*pointer).valueMethod()
	pointer.pointerMethod() // Implicitly converted to: (&value).pointerMethod()
}
