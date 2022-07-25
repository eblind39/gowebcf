package architecture

import (
	"fmt"
)

// Personas is how the archiecture package stores a person
type Persona struct {
	First string
}

// Accesor is how to store or retrieve a person
// When retriwving a person, if they do not exist, return the zero value
type Accesor interface {
	Save(n int, p Persona)
	Retrieve(n int) Persona
}

func Put(a Accesor, n int, p Persona) {
	a.Save(n, p)
}
func Get(a Accesor, n int) (Persona, error) {
	p := a.Retrieve(n)
	if p.First == "" {
		return Persona{}, fmt.Errorf("No person with n of %d", n)
	}
	return p, nil
}
