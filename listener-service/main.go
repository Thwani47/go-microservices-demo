package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"listener-service/event"
	"log"
	"math"
	"os"
	"time"
)

func main() {
	rabbitConnection, err := connect()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConnection.Close()

	log.Println("Listening for and consuming RabbitMQ messages...")
	consumer, err := event.NewConsumer(rabbitConnection)

	if err != nil {
		log.Panic(err)
	}

	err = consumer.Listen([]string{"log.info", "log.warning", "log.error"})

	if err != nil {
		log.Println(err)
	}

}

func connect() (*amqp.Connection, error) {
	var count int64
	backOff := 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			count++
		} else {
			log.Println("Connected to RabbitMQ")
			connection = c
			break
		}

		if count > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(count), 2)) * time.Second
		log.Println("backing off..")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}
