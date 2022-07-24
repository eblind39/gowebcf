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
	architecture.Put(dbm, 1, p1)
	architecture.Put(dbm, 2, p2)
	fmt.Println(architecture.Get(dbm, 1))
	fmt.Println(architecture.Get(dbm, 3))

	// or store in another data store
	architecture.Put(dbp, 1, p1)
	architecture.Put(dbp, 2, p2)
	fmt.Println(architecture.Get(dbp, 1))
	fmt.Println(architecture.Get(dbp, 2))
}
