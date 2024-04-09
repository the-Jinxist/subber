package main

import (
	"log"
	"os"
	"sync"

	"github.com/the-Jinxist/subber/config"
)

const webPort = 8080

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
	_ = config.AppConfig{
		Session:  config.GetSession(),
		Db:       config.GetDB(),
		Wait:     &wg,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	// setup mail

	// listen for web connections

}
