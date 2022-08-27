package main

import (
	"database/sql"
	"github.com/alexedwards/scs/v2"
	"github.com/eblind39/gowebcf/go-concurrency/finalproj/data"
	"log"
	"sync"
)

type Config struct {
	Session  *scs.SessionManager
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Wait     *sync.WaitGroup
	Models   data.Models
}
