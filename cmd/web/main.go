package main

import "subscription-service/configs"

const webPort = "8080"

func main() {
	// connect to the database
	db := configs.InitDB()
	db.Ping()
	// create sessions

	// create channels

	// create waitgroup

	// set up the application config

	// set up email

	// listen for web connections

}
