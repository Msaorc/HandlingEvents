package main

import (
	"fmt"

	"github.com/Msaorc/MaSystemUtils/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgs, "QueueOrders")
	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
