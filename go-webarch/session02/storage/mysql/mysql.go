package mysql

import "github.com/eblind39/gowebcf/go-webarch/session02/models"

type Db map[int]models.Car

// implement iface accesor on the db
func (m Db) Save(p int, c models.Car) {
	m[p] = c
}
func (m Db) Retrieve(p int) models.Car {
	return m[p]
}
