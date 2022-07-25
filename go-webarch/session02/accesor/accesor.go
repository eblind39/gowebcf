package accesor

import "github.com/eblind39/gowebcf/go-webarch/session02/models"

type Accesor interface {
	Save(p int, c models.Car)
	Retrieve(p int) models.Car
}

// access trough accesor
func Put(a Accesor, p int, c models.Car) {
	a.Save(p, c)
}
func Get(a Accesor, p int) models.Car {
	return a.Retrieve(p)
}
