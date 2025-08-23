# Order Processing with RabbitMQ

This demo shows how to implement a message-based order processing system using RabbitMQ in Go.

## Features

1. **Message Queuing**
   - Order submission
   - Order processing
   - Status notifications

2. **Exchange Types**
   - Direct exchange for order routing
   - Fanout exchange for notifications

3. **Message Patterns**
   - Work queues
   - Publish/Subscribe
   - Message persistence

## Architecture

```
Producer -> Orders Exchange (direct) -> [new_orders, processing_orders, completed_orders, failed_orders]
                                   -> Notifications Exchange (fanout) -> [notifications]
```

## Prerequisites

1. RabbitMQ server running locally
2. Go installed on your system

## Installation

1. Install RabbitMQ client:
   ```bash
   go get github.com/rabbitmq/amqp091-go
   ```

2. Run RabbitMQ server:
   ```bash
   # Using Docker
   docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management
   ```

3. Run the demo:
   ```bash
   # Terminal 1 - Run notification consumer
   go run . -mode consumer

   # Terminal 2 - Run order processor
   go run . -mode processor
   ```

## Components

1. **Order Processor**
   - Handles order submission
   - Routes orders through processing stages
   - Manages message persistence

2. **Notification Consumer**
   - Receives order status updates
   - Demonstrates fanout exchange pattern

3. **Message Types**
   - New orders
   - Processing orders
   - Completed orders
   - Failed orders
   - Notifications

## Best Practices Demonstrated

1. **Connection Management**
   - Proper connection handling
   - Channel management
   - Cleanup on shutdown

2. **Message Handling**
   - Message persistence
   - Proper acknowledgments
   - Error handling

3. **Exchange Design**
   - Appropriate exchange types
   - Queue bindings
   - Routing patterns

## Testing

```bash
# Start RabbitMQ
docker start rabbitmq

# Monitor RabbitMQ
open http://localhost:15672
# username: guest
# password: guest

# Run demo components
go run . -mode consumer  # Terminal 1
go run . -mode processor # Terminal 2
```

## Note

This is a demonstration project and not intended for production use. In a real application, you would want to:

- Add proper error handling
- Implement retry mechanisms
- Add dead letter queues
- Implement proper logging
- Add monitoring and metrics
- Add tests
- Handle connection recovery
- Add proper configuration management
