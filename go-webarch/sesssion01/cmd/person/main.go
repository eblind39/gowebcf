package main

import (
	"fmt"
	"github.com/eblind39/gowebcf/go-webarch/sesssion01"
	"github.com/eblind39/gowebcf/go-webarch/sesssion01/storage/mongo"
	"github.com/eblind39/gowebcf/go-webarch/sesssion01/storage/postgres"
)

func main() {
	dbm := mongo.Db{}
	dbp := postgres.Db{}

	p1 := architecture.Persona{
		First: "Juri",
	}

	p2 := architecture.Persona{
		First: "Cooper",
	}

	// Using a service
	psm := architecture.NewPersonService(dbm)
	psm.Put(1, p1)
	psm.Put(2, p2)
	fmt.Println(psm.Get(1))
	fmt.Println(psm.Get(3))

	// or store in another data store
	psp := architecture.NewPersonService(dbp)
	psp.Put(1, p1)
	psp.Put(2, p2)
	fmt.Println(psp.Get(1))
	fmt.Println(psp.Get(2))
}
