package main

import (
	"subscription-service/setup"
)

const webPort = "8080"

func main() {
	// connect to the database
	db := setup.InitDB()
	db.Ping()

	// create sessions
	sess := setup.InitSession()

	// create channels

	// create waitgroup

	// set up the application config

	// set up email

	// listen for web connections

}
