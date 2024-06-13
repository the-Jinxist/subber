package main

func (app *AppConfig) sendEmail(msg Message) {
	app.Wait.Add(1)

	app.Mailer.MailerChan <- msg
}
