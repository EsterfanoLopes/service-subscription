package main

import "subscription-service/cmd/web/mailer"

func (app *Config) sendEmail(msg mailer.Message) {
	app.Wait.Add(1)
	app.Mailer.MailerChan <- msg
}
