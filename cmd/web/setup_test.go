package main

import (
	"context"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/the-Jinxist/subber/data"
)

var testapp AppConfig

func TestMain(m *testing.M) {

	gob.Register(data.User{})

	session := scs.New()

	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	testapp = AppConfig{
		Session:       session,
		Db:            nil,
		InfoLog:       log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog:      log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
		Wait:          &sync.WaitGroup{},
	}

	// create a dummy mailer
	errorChan := make(chan error)
	mailerChan := make(chan Message, 100)
	doneChan := make(chan bool)

	testapp.Mailer = Mail{
		Wait:       testapp.Wait,
		ErrorChan:  errorChan,
		MailerChan: mailerChan,
		DoneChan:   doneChan,
	}

	go func() {
		select {
		case <-testapp.Mailer.MailerChan:
		case <-testapp.Mailer.ErrorChan:
		case <-testapp.Mailer.DoneChan:
			return
		}

	}()

	go func() {
		for {
			select {
			case err := <-testapp.ErrorChan:
				testapp.ErrorLog.Println(err)
			case <-testapp.ErrorChanDone:
				return
			}
		}
	}()

	os.Exit(m.Run())
}

func getCtx(req *http.Request) context.Context {
	ctx, err := testapp.Session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}

	return ctx

}
