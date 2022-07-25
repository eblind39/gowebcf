package sqlserver

import "github.com/eblind39/gowebcf/go-webarch/session02/models"

type Db map[int]models.Car

// implement iface accesor on the db
func (s Db) Save(p int, c models.Car) {
	s[p] = c
}
func (s Db) Retrieve(p int) models.Car {
	return s[p]
}
