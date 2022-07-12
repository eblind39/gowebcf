package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age int
}

type ByAge []Person
func (a ByAge) Len() int { return len((a)) }
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person
func (a ByName) Len() int { return len((a)) }
func (a ByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }


func main() {
	p1 := Person{"Eduar", 32}
	p2 := Person{"Mary", 17}
	p3 := Person{"Caroline", 28}
	p4 := Person{"Adriana", 35}

	persons :=  []Person{p1, p2, p3, p4}

	sort.Sort(ByAge(persons))
	fmt.Println(persons)

	sort.Sort(ByName(persons))
	fmt.Println(persons)
}