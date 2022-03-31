package handler

import (
	"encoding/json"
	"fmt"
	"gochicoba/consumer"
	"gochicoba/models"
	"gochicoba/repository"
	"log"

	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func BuyBrokerListener(db *gorm.DB, amqp *amqp.Connection) {
	channel := consumer.InitChannelMessageBroker(amqp)

	res, err := channel.Consume(
		"QueueService1", // queue name
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no local
		false,           // no wait
		nil,             // arguments
	)

	if err != nil {
		log.Println(err)
	}

	// make a channel to receive messages into infinite loop
	forever := make(chan bool)

	go func() {
		for message := range res {
			// for example, show received message in a console
			log.Printf(" > Received message: %s\n", message.Body)
			var mT *models.Transaction = new(models.Transaction)
			json.Unmarshal(message.Body, mT)
			fmt.Println(mT)
			repository.CreateTransactionBroker(db, mT)
		}
	}()

	<-forever
}
