# Custom Network Protocol Demo

This demo implements a custom network protocol with Go, demonstrating various aspects of network programming.

## Features

1. Custom Protocol Implementation
   - Version control
   - Command types
   - Checksum validation
   - Frame marshaling/unmarshaling

2. Server Implementation
   - Connection pooling
   - Metrics collection (Prometheus)
   - Graceful shutdown
   - Error handling

3. Client Implementation
   - Ping/Pong mechanism
   - Message sending
   - File transfer
   - Connection management

## Protocol Specification

### Frame Format
```
+----------------+----------------+----------------+----------------+
|    Version     |    Command    |        Payload Length         |
+----------------+----------------+----------------+----------------+
|                           Payload...                           |
+----------------+----------------+----------------+----------------+
|                           Checksum                            |
+----------------+----------------+----------------+----------------+
```

### Commands
1. Ping (1): Simple ping/pong mechanism
2. Message (2): Text message transfer
3. File (3): File transfer

## Usage

### Building
```bash
go build -o server ./cmd/server
go build -o client ./cmd/client
```

### Running the Server
```bash
./server -addr :8080 -pool-size 10
```

### Running the Client
```bash
# Send a message
./client -addr localhost:8080 -message "Hello, World!"

# Send a file
./client -addr localhost:8080 -file path/to/file.txt
```

## Implementation Details

### Server Features
- Connection pooling for resource management
- Prometheus metrics integration
- Graceful shutdown handling
- Concurrent client handling

### Client Features
- Automatic reconnection
- File transfer support
- Ping/Pong health check
- Command-line interface

### Protocol Features
- Version checking
- Checksum validation
- Maximum payload size enforcement
- Command validation

## Metrics

The server exposes the following Prometheus metrics:
- active_connections
- total_messages
- message_latency_seconds

## Best Practices Demonstrated

1. Error Handling
   - Network errors
   - Protocol errors
   - Resource cleanup

2. Performance
   - Connection pooling
   - Buffer management
   - Concurrent processing

3. Monitoring
   - Prometheus integration
   - Connection tracking
   - Message latency

4. Security
   - Checksum validation
   - Protocol version control
   - Size limitations

## Requirements

- Go 1.21 or later
- Prometheus client library
