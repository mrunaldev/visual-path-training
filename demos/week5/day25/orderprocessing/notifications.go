package main

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// NotificationConsumer handles order notifications
func NotificationConsumer(amqpURI string) error {
	// Connect to RabbitMQ
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create channel
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Declare notification queue
	q, err := ch.QueueDeclare(
		"notifications", // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		return err
	}

	// Bind to notifications exchange
	err = ch.QueueBind(
		q.Name,          // queue name
		"",              // routing key
		"notifications", // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Consume messages
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	log.Println("Notification consumer started. Waiting for messages...")

	for msg := range msgs {
		var order Order
		if err := json.Unmarshal(msg.Body, &order); err != nil {
			log.Printf("Error unmarshaling notification: %v", err)
			continue
		}

		log.Printf("Notification: Order %s (Customer: %s) status changed to %s",
			order.ID, order.CustomerID, order.Status)
	}

	return nil
}
