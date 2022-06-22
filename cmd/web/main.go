package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"subscription-service/data"
	"subscription-service/setup"
	"sync"
	"syscall"
)

const webPort = "8080"

func main() {
	// create channels

	// create waitgroup
	var wg sync.WaitGroup

	// init db
	db := setup.InitDB()

	// set up the application config
	app := Config{
		Session:  setup.InitSession(),
		DB:       db,
		InfoLog:  setup.NewInfoLogger(),
		ErrorLog: setup.NewErrorLogger(),
		Wait:     &wg,
		Models:   data.New(db),
	}

	// set up email

	// listen for signals
	go app.ListenForShutdown()

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

func (app *Config) ListenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.shutdown()
	os.Exit(0)
}

func (app *Config) shutdown() {
	// perform any cleanup tasks
	app.InfoLog.Println("would run cleanup tasks...")

	// block until witgroup is empty
	app.Wait.Wait()

	app.InfoLog.Println("closing channels and shutting down application...")
}
