package client

import (
	"encoding/json"
	. "github.com/Andrew161644/currency_exchange_grpc/api/model"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func ExchangerRPC(task CurrencyExchangeTask, host string) (res CurrencyExchangeTask, err error) {
	conn, err := amqp.Dial(host)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	corrId := randomString(32)
	body, err := json.Marshal(task)
	err = ch.Publish(
		"",          // exchange
		"exchange", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          body,
		})
	failOnError(err, "Failed to publish a message")



	for d := range msgs {
		if corrId == d.CorrelationId {
			task := &CurrencyExchangeTask{}
			err := json.Unmarshal(d.Body, task)
			if err!=nil {
				return CurrencyExchangeTask{}, err
			}
			return *task, nil
		}
	}

	return
}