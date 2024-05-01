package main

import (
	"database/sql"
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"
	"github.com/the-Jinxist/subber/data"
)

type AppConfig struct {
	Session  *scs.SessionManager
	Db       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Wait     *sync.WaitGroup
	Models   data.Models
}
