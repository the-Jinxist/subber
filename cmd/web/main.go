package main

import (
	"github.com/the-Jinxist/subber/config"
)

const webPort = 8080

func main() {

	// connect to db
	config.LoadEnvs("../..")

	config.InitDB()

	// create sessions

	// create channels

	// create waitgroup

	//set up the application config

	// setup mail

	// listen for web connections

}
