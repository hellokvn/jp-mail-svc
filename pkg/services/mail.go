package services

import (
	"fmt"

	"github.com/hellokvn/jp-mail-svc/pkg/db"
)

type Server struct {
	H db.Handler
}

type SendMailBody struct {
	Template string `json:"template"`
	To       string `json:"to"`
}

func (s *Server) SendMail(b SendMailBody) {
	fmt.Println("SendMail Template", b.Template)
	fmt.Println("SendMail To", b.To)
}
