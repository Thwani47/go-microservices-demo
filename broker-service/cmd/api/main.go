package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"math"
	"net/http"
	"os"
	"time"
)

const webPort = "80"

type Config struct {
	RabbitMQ *amqp.Connection
}

func main() {
	rabbitConnection, err := connect()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConnection.Close()

	app := Config{
		RabbitMQ: rabbitConnection,
	}
	log.Printf("Starting broker service on port %s\n", webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start server
	err = srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
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
