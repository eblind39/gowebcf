package main

import (
	"fmt"
	"github.com/eblind39/gowebcf/go-webarch/session02/accesor"
	"github.com/eblind39/gowebcf/go-webarch/session02/models"
	"github.com/eblind39/gowebcf/go-webarch/session02/storage/mysql"
	"github.com/eblind39/gowebcf/go-webarch/session02/storage/sqlserver"
)

func main() {
	c1 := models.Car{
		Make: "Toyota",
	}
	c2 := models.Car{
		Make: "Jeep",
	}

	msql := mysql.Db{}
	sqls := sqlserver.Db{}

	accesor.Put(msql, 1, c1)
	accesor.Put(sqls, 1, c2)

	fmt.Println("car from MySQL", accesor.Get(msql, 1))
	fmt.Println("car from SQLServer", accesor.Get(sqls, 1))
}
