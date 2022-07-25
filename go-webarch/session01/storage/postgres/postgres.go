package postgres

import (
	"github.com/eblind39/gowebcf/go-webarch/session01"
)

type Db map[int]architecture.Persona

func (pg Db) Save(n int, p architecture.Persona) {
	pg[n] = p
}
func (pg Db) Retrieve(n int) (p architecture.Persona) {
	return pg[n]
}
