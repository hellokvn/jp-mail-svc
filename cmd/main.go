package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hellokvn/jp-mail-svc/pkg/config"
	"github.com/hellokvn/jp-mail-svc/pkg/db"
	"github.com/hellokvn/jp-mail-svc/pkg/services"
	"github.com/streadway/amqp"
)

type Message struct {
	Template string `json:"template"`
	To       string `json:"to"`
}

func main() {
	config, err := config.LoadConfig()

	h := db.Init(config.DBUrl)
	s := services.Server{
		H: h,
	}

	connectRabbitMQ, err := amqp.Dial(config.AMQPUrl)

	if err != nil {
		panic(err)
	}

	defer connectRabbitMQ.Close()

	channelRabbitMQ, err := connectRabbitMQ.Channel()

	if err != nil {
		panic(err)
	}

	defer channelRabbitMQ.Close()

	// Subscribing to QueueService1 for getting messages.
	messages, err := channelRabbitMQ.Consume(
		"QueueService1",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println(err)
	}

	log.Println("Successfully connected to RabbitMQ")

	forever := make(chan bool)

	go func() {
		for message := range messages {
			var b services.SendMailBody
			err := json.Unmarshal(message.Body, &b)

			if err != nil {
				fmt.Println(err)
			}

			switch message.Type {
			case "send_mail":
				s.SendMail(b)
			default:
				fmt.Println("Invalid message type")
			}
		}
	}()

	<-forever
}
