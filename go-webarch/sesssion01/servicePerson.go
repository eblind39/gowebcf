package architecture

import (
	"fmt"
)

type Persona struct {
	First string
}

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
