package main

import (
	"database/sql"
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"
	"github.com/the-Jinxist/subber/data"
)

type AppConfig struct {
	Session       *scs.SessionManager
	Db            *sql.DB
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Wait          *sync.WaitGroup
	Models        data.Models
	ErrorChan     chan error
	ErrorChanDone chan bool

	Mailer Mail
}

func (app *AppConfig) createMail() Mail {

	errChan := make(chan error)
	mailerChan := make(chan Message, 100)
	mailerDone := make(chan bool)

	m := Mail{
		Domain:      "localhost",
		Host:        "localhost",
		Port:        1025,
		Encryption:  "none",
		FromAddress: "info@mycompany.com",
		FromName:    "info",
		ErrorChan:   errChan,
		MailerChan:  mailerChan,
		DoneChan:    mailerDone,
		Wait:        app.Wait,
	}

	return m

}
