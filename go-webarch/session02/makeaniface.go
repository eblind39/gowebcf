package main

import "fmt"

type milliseconds int

type computer interface {
	start() milliseconds
}

type desktop struct {
	make string
}

type laptop struct {
	make string
}

func (d desktop) start() milliseconds {
	return 750
}

func (l laptop) start() milliseconds {
	return 250
}

func main() {
	d := desktop{
		make: "Dell",
	}
	fmt.Println(d.make, " starts in ", d.start(), "ms")

	l := laptop{
		make: "Allienware",
	}
	fmt.Println(l.make, " starts in ", l.start(), "ms")

	iface1 := d
	iface2 := l
	fmt.Printf("%T\n", iface1)
	fmt.Printf("%T\n", iface2)
	fmt.Println(iface1.make, " starts in ", iface1.start(), "ms")
	fmt.Println(iface2.make, " starts in ", iface2.start(), "ms")
}
