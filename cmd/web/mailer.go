package main

import (
	"sync"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

type Mail struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string

	Wait       *sync.WaitGroup
	MailerChan chan Message
	ErrorChan  chan error
	DoneChan   chan bool
}

type Message struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Attachments []string
	Data        any
	DataMap     map[string]any
	Template    string
}

// function to listen for messages on mailer chan

func (m *Mail) SendEmail(msg Message, erroChan chan error) {
	if msg.Template == "" {
		msg.Template = "mail"
	}

	if msg.From == "" {
		msg.From = m.FromAddress
	}

	if msg.FromName == "" {
		msg.FromName = m.FromName
	}

	data := map[string]any{
		"message": msg.Data,
	}
	msg.DataMap = data

	//build html message
	formattedMessage, err := m.buildHTMLMessage(msg)
	if err != nil {
		erroChan <- err
	}

	plainTextMessage, err := m.buildPlainTextMessage(msg)
	if err != nil {
		erroChan <- err
	}

	server := mail.NewSMTPClient()
	server.Host = m.Host
	server.Port = m.Port
	server.Username = m.Username
	server.Password = m.Password
	server.Encryption = m.getEncryption()
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		erroChan <- err
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).AddTo(msg.To).SetSubject(msg.Subject)
	email.SetBody(mail.TextPlain, plainTextMessage)
	email.AddAlternative(mail.TextHTML, formattedMessage)

	if len(msg.Attachments) > 0 {
		for _, v := range msg.Attachments {
			email.AddAttachment(v)
		}
	}

	err = email.Send(smtpClient)
	if err != nil {
		erroChan <- err
	}

}

func (m *Mail) buildHTMLMessage(msg Message) (string, error) {
	return "", nil
}

func (m *Mail) buildPlainTextMessage(msg Message) (string, error) {
	return "", nil
}

func (m *Mail) getEncryption() mail.Encryption {
	switch m.Encryption {
	case "tls":
		{
			return mail.EncryptionSTARTTLS
		}
	case "ssl":
		{
			return mail.EncryptionSSL
		}
	case "none":
		{
			return mail.EncryptionNone
		}
	default:
		{
			return mail.EncryptionSTARTTLS
		}
	}
}
