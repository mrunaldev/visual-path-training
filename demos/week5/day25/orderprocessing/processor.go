package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Order represents a customer order
type Order struct {
	ID          string    `json:"id"`
	CustomerID  string    `json:"customer_id"`
	Items       []Item    `json:"items"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

// Item represents an order item
type Item struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

// OrderProcessor handles order processing through RabbitMQ
type OrderProcessor struct {
	conn    *amqp.Connection
	ch      *amqp.Channel
	ctx     context.Context
	cancel  context.CancelFunc
	cleanup func()
}

// NewOrderProcessor creates a new order processor
func NewOrderProcessor(amqpURI string) (*OrderProcessor, error) {
	// Create context with cancellation
	ctx, cancel := context.WithCancel(context.Background())

	// Connect to RabbitMQ
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		cancel()
		return nil, err
	}

	// Create channel
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		cancel()
		return nil, err
	}

	// Create cleanup function
	cleanup := func() {
		ch.Close()
		conn.Close()
		cancel()
	}

	processor := &OrderProcessor{
		conn:    conn,
		ch:      ch,
		ctx:     ctx,
		cancel:  cancel,
		cleanup: cleanup,
	}

	// Set up exchanges and queues
	if err := processor.setup(); err != nil {
		cleanup()
		return nil, err
	}

	return processor, nil
}

// setup creates exchanges and queues
func (p *OrderProcessor) setup() error {
	// Declare exchanges
	exchanges := []struct {
		name string
		kind string
	}{
		{"orders", "direct"},
		{"notifications", "fanout"},
	}

	for _, e := range exchanges {
		err := p.ch.ExchangeDeclare(
			e.name, // name
			e.kind, // type
			true,   // durable
			false,  // auto-deleted
			false,  // internal
			false,  // no-wait
			nil,    // arguments
		)
		if err != nil {
			return err
		}
	}

	// Declare queues
	queues := []struct {
		name       string
		routingKey string
		exchange   string
	}{
		{"new_orders", "new", "orders"},
		{"processing_orders", "processing", "orders"},
		{"completed_orders", "completed", "orders"},
		{"failed_orders", "failed", "orders"},
		{"notifications", "", "notifications"},
	}

	for _, q := range queues {
		_, err := p.ch.QueueDeclare(
			q.name, // name
			true,   // durable
			false,  // delete when unused
			false,  // exclusive
			false,  // no-wait
			nil,    // arguments
		)
		if err != nil {
			return err
		}

		// Bind queue to exchange if routing key is provided
		if q.exchange != "" {
			err = p.ch.QueueBind(
				q.name,       // queue name
				q.routingKey, // routing key
				q.exchange,   // exchange
				false,
				nil,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// SubmitOrder submits a new order for processing
func (p *OrderProcessor) SubmitOrder(order Order) error {
	order.Status = "new"
	order.CreatedAt = time.Now()

	data, err := json.Marshal(order)
	if err != nil {
		return err
	}

	return p.ch.PublishWithContext(p.ctx,
		"orders", // exchange
		"new",    // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent,
		},
	)
}

// ProcessOrders starts processing orders
func (p *OrderProcessor) ProcessOrders() error {
	msgs, err := p.ch.Consume(
		"new_orders", // queue
		"",           // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		var order Order
		if err := json.Unmarshal(msg.Body, &order); err != nil {
			log.Printf("Error unmarshaling order: %v", err)
			msg.Reject(false)
			continue
		}

		// Process the order (simulated)
		log.Printf("Processing order %s for customer %s", order.ID, order.CustomerID)
		time.Sleep(time.Second) // Simulate processing

		// Update order status
		order.Status = "processing"

		// Publish to processing queue
		data, _ := json.Marshal(order)
		err = p.ch.PublishWithContext(p.ctx,
			"orders",     // exchange
			"processing", // routing key
			false, false,
			amqp.Publishing{
				ContentType:  "application/json",
				Body:         data,
				DeliveryMode: amqp.Persistent,
			},
		)
		if err != nil {
			msg.Reject(false)
			continue
		}

		// Notify all interested parties
		p.ch.PublishWithContext(p.ctx,
			"notifications", // exchange
			"",              // routing key
			false, false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        data,
			},
		)

		msg.Ack(false)
	}

	return nil
}

// Close closes the order processor
func (p *OrderProcessor) Close() {
	p.cleanup()
}
