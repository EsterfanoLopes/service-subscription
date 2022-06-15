package main

import (
	"subscription-service/setup"
	"sync"
)

const webPort = "8080"

func main() {
	// create channels

	// create waitgroup
	var wg sync.WaitGroup

	// set up the application config
	app := Config{
		Session:  setup.InitSession(),
		DB:       setup.InitDB(),
		InfoLog:  setup.NewInfoLogger(),
		ErrorLog: setup.NewErrorLogger(),
		Wait:     &wg,
	}

	// set up email

	// listen for web connections

}
