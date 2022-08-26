package main

import (
	"database/sql"
	"github.com/alexedwards/scs/v2"
	"log"
	"sync"
)

type Config struct {
	Session  *scs.SessionManager
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Wait     *sync.WaitGroup
}
