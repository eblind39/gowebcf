package main

import "fmt"

type car struct {
	make string
}

func (c car) String() string {
	return fmt.Sprintf("My make is: %s", c.make)
}

func main() {
	c1 := car{
		make: "Honda",
	}

	fmt.Println(c1)

	c2 := car{
		make: "Toyota",
	}

	m := map[car]string{}
	m[c1] = "loved in all over the world"
	m[c2] = "most popular"

	fmt.Println(c1, " - ", m[c1])
	fmt.Println(c2, " - ", m[c2])
	// do not use the following one if we are using pointers to struct
	fmt.Println(c1 == c2)
	fmt.Println(c1 == c1)
}
