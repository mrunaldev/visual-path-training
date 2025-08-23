---
marp: true
theme: default
paginate: true
---

# Message Queuing with RabbitMQ
## Advanced Go Programming - Day 25

---

# Overview

1. Introduction to Message Queuing
2. RabbitMQ Concepts
3. AMQP Protocol
4. Go-AMQP Library
5. Common Patterns
6. Best Practices

---

# Message Queuing Benefits

- Decoupling services
- Asynchronous communication
- Load balancing
- Scalability
- Fault tolerance
- Message persistence

---

# RabbitMQ Concepts

1. Producer
   - Sends messages to exchange

2. Exchange
   - Routes messages to queues
   - Different types (direct, fanout, topic, headers)

3. Queue
   - Stores messages
   - FIFO order

4. Consumer
   - Receives messages from queue

---

# Exchange Types

1. **Direct Exchange**
   ```
   Producer -> Exchange [routing key] -> Queue
   ```

2. **Fanout Exchange**
   ```
   Producer -> Exchange -> All bound queues
   ```

3. **Topic Exchange**
   ```
   Producer -> Exchange [pattern matching] -> Matching queues
   ```

4. **Headers Exchange**
   ```
   Producer -> Exchange [header attributes] -> Matching queues
   ```

---

# AMQP in Go

```go
import "github.com/rabbitmq/amqp091-go"

// Connect to RabbitMQ
conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

// Create a channel
ch, err := conn.Channel()

// Declare a queue
q, err := ch.QueueDeclare(
    "hello", // name
    false,   // durable
    false,   // delete when unused
    false,   // exclusive
    false,   // no-wait
    nil,     // arguments
)
```

---

# Publishing Messages

```go
// Publish a message
err = ch.PublishWithContext(ctx,
    "",     // exchange
    q.Name, // routing key
    false,  // mandatory
    false,  // immediate
    amqp.Publishing{
        ContentType: "text/plain",
        Body:        []byte("Hello World!"),
    })
```

---

# Consuming Messages

```go
msgs, err := ch.Consume(
    q.Name, // queue
    "",     // consumer
    true,   // auto-ack
    false,  // exclusive
    false,  // no-local
    false,  // no-wait
    nil,    // args
)

for msg := range msgs {
    fmt.Printf("Received: %s\n", msg.Body)
}
```

---

# Message Acknowledgment

```go
msgs, err := ch.Consume(
    q.Name, // queue
    "",     // consumer
    false,  // auto-ack
    false,  // exclusive
    false,  // no-local
    false,  // no-wait
    nil,    // args
)

for msg := range msgs {
    // Process message
    
    // Manual acknowledgment
    msg.Ack(false) // false = ack single message
}
```

---

# Exchange Declaration

```go
err = ch.ExchangeDeclare(
    "logs",   // name
    "fanout", // type
    true,     // durable
    false,    // auto-deleted
    false,    // internal
    false,    // no-wait
    nil,      // arguments
)
```

---

# Queue Binding

```go
err = ch.QueueBind(
    q.Name,    // queue name
    "black.*", // routing key
    "logs",    // exchange
    false,     // no-wait
    nil,       // arguments
)
```

---

# Publish/Subscribe Pattern

```go
// Publisher
err = ch.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)
err = ch.PublishWithContext(ctx,
    "logs", // exchange
    "",     // routing key
    false,  // mandatory
    false,  // immediate
    amqp.Publishing{Body: []byte(message)})

// Subscriber
q, err := ch.QueueDeclare("", false, false, true, false, nil)
err = ch.QueueBind(q.Name, "", "logs", false, nil)
msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
```

---

# Routing Pattern

```go
// Publisher
err = ch.ExchangeDeclare("logs_direct", "direct", true, false, false, false, nil)
err = ch.PublishWithContext(ctx,
    "logs_direct", // exchange
    severity,      // routing key
    false, false,
    amqp.Publishing{Body: []byte(message)})

// Subscriber
err = ch.QueueBind(q.Name, severity, "logs_direct", false, nil)
```

---

# Topic Pattern

```go
// Publisher
err = ch.ExchangeDeclare("logs_topic", "topic", true, false, false, false, nil)
err = ch.PublishWithContext(ctx,
    "logs_topic",     // exchange
    "kern.critical",  // routing key
    false, false,
    amqp.Publishing{Body: []byte(message)})

// Subscriber
err = ch.QueueBind(q.Name, "kern.*", "logs_topic", false, nil)
```

---

# Message Persistence

```go
err = ch.PublishWithContext(ctx,
    "",
    q.Name,
    false,
    false,
    amqp.Publishing{
        DeliveryMode: amqp.Persistent,
        ContentType:  "text/plain",
        Body:        []byte(message),
    })
```

---

# Error Handling

```go
// Connection monitoring
go func() {
    for {
        reason, ok := <-conn.NotifyClose(make(chan *amqp.Error))
        if !ok {
            return
        }
        log.Printf("Connection closed: %s", reason)
        // Implement reconnection logic
    }
}()

// Channel monitoring
go func() {
    for {
        reason, ok := <-ch.NotifyClose(make(chan *amqp.Error))
        if !ok {
            return
        }
        log.Printf("Channel closed: %s", reason)
        // Implement channel recovery
    }
}()
```

---

# Best Practices

1. Connection Management
   - Connection pooling
   - Automatic reconnection
   - Channel management

2. Message Handling
   - Proper acknowledgment
   - Dead letter queues
   - Message TTL

3. Error Handling
   - Retry mechanisms
   - Circuit breakers
   - Monitoring

4. Performance
   - Channel sharing
   - Prefetch settings
   - Batch processing

---

# Common Use Cases

1. Task Distribution
2. Real-time Updates
3. Event-driven Architecture
4. Log Aggregation
5. Distributed Systems
6. Microservices Communication

---

# Questions?

1. When to use RabbitMQ?
2. How to handle failures?
3. Best practices for scaling?
4. Message delivery guarantees?
