# Go Microservices Demo

Building small, self-contained microservices that communicate with one another and a simple front-end application over:
- a REST API
- RPC (Remote Procedure Call)
- gRPC (Google RPC)
- AMPQ (Advanced Message Queueing Protocol)

The microservices include (all written in Go):
- A **front-end** service that displays web pages
- An **authentication** service that authenticates users in a Postgres database
- A **logging** service that logs data to a MongoDB database
- A **listerner** service that received messages from RabbitMQ and acts upon them
- A **broker** service which is the single entry point into the microservices cluster
- A **mail** service, which takes a JSON payload, converts it into a formatted email, and sends it out

This is a code-along project to this Udemy Course: [Working with Microservices in Go (Golang)](https://www.udemy.com/course/working-with-microservices-in-go/)

# Running project
## Using make
```bash
make up_build
```