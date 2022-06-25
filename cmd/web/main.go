package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"subscription-service/cmd/web/mailer"
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
	app.Mailer = app.createMail()
	go app.ListenForMail()

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

	app.Mailer.DoneChan <- true

	app.InfoLog.Println("closing channels and shutting down application...")
	close(app.Mailer.MailerChan)
	close(app.Mailer.ErrorChan)
	close(app.Mailer.DoneChan)
}

func (app *Config) createMail() mailer.Mail {
	errorChan := make(chan error)
	mailerChan := make(chan mailer.Message, 100)
	mailerDoneChan := make(chan bool)

	m := mailer.Mail{
		Domain:      "localhost",
		Host:        "localhost",
		Port:        1025,
		Encryption:  "none",
		FromAddress: "info@test.com",
		FromName:    "info",
		Wait:        app.Wait,
		ErrorChan:   errorChan,
		MailerChan:  mailerChan,
		DoneChan:    mailerDoneChan,
	}

	return m
}
