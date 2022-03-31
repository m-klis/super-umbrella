package consumer

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func InitializeMessageBroker() *amqp.Connection {
	amqpHost := os.Getenv("AMQP_HOST")
	amqpUser := os.Getenv("AMQP_USER")
	amqpPass := os.Getenv("AMQP_PASS")
	amqpPort := os.Getenv("AMQP_PORT")
	link := "amqp://" + amqpUser + ":" + amqpPass + "@" + amqpHost + ":" + amqpPort + "/"

	ConnectRabbitMQ, err := amqp.Dial(link)
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to RabbitMQ")
	return ConnectRabbitMQ
}

func InitChannelMessageBroker(ac *amqp.Connection) *amqp.Channel {
	channelRabbitMQ, err := ac.Channel()
	if err != nil {
		panic(err)
	}
	log.Println("Waiting for messages")
	return channelRabbitMQ
}

func CloseMessageBroker(ac *amqp.Connection) error {
	err := ac.Close()
	if err != nil {
		return err
	}
	return nil
}
