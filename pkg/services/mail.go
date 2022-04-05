package services

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/hellokvn/jp-mail-svc/pkg/config"
)

type Server struct {
	C config.Config
}

type SendMailBody struct {
	Template string `json:"template"`
	To       string `json:"to"`
}

func (s *Server) SendMail(b *SendMailBody) {
	to := []string{b.To}
	addr := fmt.Sprintf("%s:%s", s.C.MailHost, s.C.MailPort)

	// keeping it simple
	msg := []byte("From: " + s.C.MailFrom + "\r\n" +
		"To: " + b.To + "\r\n" +
		"Subject: Registered\r\n\r\n" +
		"Successfully registered\r\n")

	auth := smtp.PlainAuth("", s.C.MailUser, s.C.MailPassword, s.C.MailHost)
	err := smtp.SendMail(addr, auth, s.C.MailFrom, to, msg)

	if err != nil {
		log.Fatal(err)
	}
}
