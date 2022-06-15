package main

import (
	"fmt"
	"log"
	"net/http"
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
	app.serve()
}

func (app *Config) serve() {
	// start http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	app.InfoLog.Println("Starting web server...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
