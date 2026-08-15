package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/finkord/building_modern_web_app_with_go/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
			sendMsg(msg)
		}
	}()
}

func sendMsg(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
		return
	}
	email := mail.NewMSG()
	email.SetFrom(m.From)
	email.AddTo(m.To)
	email.SetSubject(m.Subject)
	if m.Template == "" {
		email.SetBody(mail.TextHTML, m.Content)
	} else {
		// data, err := ioutil.ReadFile(fmt.Sprintf("./email-templates/%s", m.Template))
		data, err := os.ReadFile("./email-templates/" + m.Template)
		if err != nil {
			errorLog.Println(err)
			return
		}
		mailTemplate := string(data)
		msgToSend := strings.Replace(mailTemplate, "[%body%]", m.Content, 1)
		email.SetBody(mail.TextHTML, msgToSend)
	}

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Email sent successfully to %s\n", m.To)
	}
}
