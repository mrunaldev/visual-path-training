# Network Programming in Go

This directory contains examples of network programming in Go, demonstrating both TCP and UDP communication.

## Structure

```
.
├── tcp/
│   ├── server/      # TCP server implementation
│   └── client/      # TCP client implementation
└── udp/
    ├── server/      # UDP server implementation
    └── client/      # UDP client implementation
```

## Running the Examples

### TCP Example

1. First, start the TCP server:
   ```
   cd tcp/server
   go run main.go
   ```

2. In another terminal, start the TCP client:
   ```
   cd tcp/client
   go run main.go
   ```

3. Type messages in the client terminal and see them echoed back by the server.

### UDP Example

1. First, start the UDP server:
   ```
   cd udp/server
   go run main.go
   ```

2. In another terminal, start the UDP client:
   ```
   cd udp/client
   go run main.go
   ```

3. Type messages in the client terminal and see them echoed back by the server.

## Key Features

- TCP and UDP implementations
- Concurrent handling of multiple clients (TCP)
- Simple echo server functionality
- Interactive client interface
- Proper error handling
- Clean shutdown with deferred connections

## Learning Objectives

1. Understanding TCP vs UDP
2. Working with network connections in Go
3. Handling concurrent connections
4. Proper error handling in network code
5. Using Go's net package effectively
