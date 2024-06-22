package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/the-Jinxist/subber/config"
	"github.com/the-Jinxist/subber/data"
)

const webPort = "80"

func main() {

	// connect to db
	config.LoadEnvs("../..")

	config.InitDB()

	// create sessions
	config.InitRedis()
	config.InitSession()

	//create logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// create channels

	// create waitgroup
	wg := sync.WaitGroup{}

	//set up the application config
	appConfig := AppConfig{
		Session:       config.GetSession(),
		Db:            config.GetDB(),
		Wait:          &wg,
		InfoLog:       infoLog,
		ErrorLog:      errorLog,
		Models:        data.New(config.GetDB()),
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
	}

	// setup mail
	appConfig.Mailer = appConfig.createMail()
	go appConfig.ListenForMail()

	// listen for signals
	go appConfig.listenForShutdown()

	// listen for errors
	go appConfig.listenForErrors()

	// listen for web connections
	appConfig.serve()

}

func (app *AppConfig) serve() {
	//start http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	app.InfoLog.Println("starting web server")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic("error while creating server; %w", err)
	}

}

func (app *AppConfig) listenForShutdown() {
	quit := make(chan os.Signal, 1)

	//this listens for the signal to interrupt or terminate the channel
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	//this blocks until we get a value in the quit channel
	<-quit

	app.shutdown()

	os.Exit(0)

}

func (app *AppConfig) listenForErrors() {
	for {
		select {
		case err := <-app.ErrorChan:
			app.ErrorLog.Println(err)
		case <-app.ErrorChanDone:
			return
		}
	}
}

func (app *AppConfig) shutdown() {
	//perform any cleanup task
	app.InfoLog.Println("would run clean up task")

	//block until waitgroup is empty
	app.Wait.Wait()
	app.Mailer.DoneChan <- true
	app.ErrorChanDone <- true

	app.InfoLog.Println("closing channels and shutting down application")
	close(app.Mailer.MailerChan)
	close(app.Mailer.ErrorChan)
	close(app.Mailer.DoneChan)

	close(app.ErrorChan)
	close(app.ErrorChanDone)

}
