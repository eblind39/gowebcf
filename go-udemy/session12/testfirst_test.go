package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestMiSuma(t *testing.T) {
	// table of tests
	type prueba struct {
		datos     []int
		respuesta int
	}
	tabletest := []prueba{
		prueba{datos: []int{2, 4, 6}, respuesta: 12},
		prueba{datos: []int{1, 3, 5}, respuesta: 9},
		prueba{datos: []int{3, 5}, respuesta: 8},
		prueba{datos: []int{7, 1, 3}, respuesta: 11},
		prueba{datos: []int{-10, 4, 20}, respuesta: 14},
	}

	v := miSuma(2, 8)
	if v != 10 {
		t.Error("Expected", 10, "Got", miSuma(2, 8))
	}

	for _, pr := range tabletest {
		v = miSuma(pr.datos...)
		if v != pr.respuesta {
			t.Error("Expected", pr.respuesta, "Got", v)
		}
	}
}

func BenchmarkMiSuma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := miSuma(2, 5, 4, 3, 7, 6, 1)
		fmt.Println(v)
	}
}

func TestCat(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = Cat(xs)
	if s != "Shaken not stirred" {
		t.Error("got", s, "want", "Shaken not stirred")
	}
}

func TestJoin(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = Join(xs)
	if s != "Shaken not stirred" {
		t.Error("got", s, "want", "Shaken not stirred")
	}
}

func ExampleCat() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Cat(xs))
	// Output:
	// Shaken not stirred
}

func ExampleJoin() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs))
	// Output:
	// Shaken not stirred
}

const s = "We ask ourselves, Who am I to be brilliant, gorgeous, talented, fabulous? Actually, who are you not to be? Your playing small does not serve the world. There is nothing enlightened about shrinking so that other people won't feel insecure around you. We are all meant to shine, as children do. We were born to make manifest the glory that is within us. It's not just in some of us; it's in everyone. And as we let our own light shine, we unconsciously give other people permission to do the same. As we are liberated from our own fear, our presence automatically liberates others. - Marianne Williamson"

var xs []string

func BenchmarkCat(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}

func BenchmarkJoin(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}
