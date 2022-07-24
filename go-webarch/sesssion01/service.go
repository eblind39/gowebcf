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

type PersonService struct {
	a Accesor
}

func NewPersonService(a Accesor) PersonService {
	return PersonService{
		a: a,
	}
}

func (ps PersonService) Put(n int, p Persona) {
	ps.a.Save(n, p)
}
func (ps PersonService) Get(n int) (Persona, error) {
	p := ps.a.Retrieve(n)
	if p.First == "" {
		return Persona{}, fmt.Errorf("No person with n of %d", n)
	}
	return p, nil
}

func PutAcc(a Accesor, n int, p Persona) {
	a.Save(n, p)
}
func GetAcc(a Accesor, n int) (Persona, error) {
	p := a.Retrieve(n)
	if p.First == "" {
		return Persona{}, fmt.Errorf("No person with n of %d", n)
	}
	return p, nil
}
