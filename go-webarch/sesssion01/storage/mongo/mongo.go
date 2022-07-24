package mongo

import (
	"github.com/eblind39/gowebcf/go-webarch/sesssion01"
)

type Db map[int]architecture.Persona

func (m Db) Save(n int, p architecture.Persona) {
	m[n] = p
}
func (m Db) Retrieve(n int) (p architecture.Persona) {
	return m[n]
}
