---
marp: true
theme: default
paginate: true
---

# Network Programming in Go
## Week 4 - Day 19

---

# Today's Topics

1. Network Basics
2. TCP Programming
3. UDP Programming
4. Connection Handling
5. Error Management

---

# Network Basics

```go
// Network types in Go
net.IP          // IP address
net.IPNet       // IP network
net.Interface   // Network interface
net.Addr        // Network address
```

- TCP vs UDP
- IPv4 vs IPv6
- Ports and Protocols

---

# TCP Server Example

```go
listener, err := net.Listen("tcp", ":8080")
if err != nil {
    log.Fatal(err)
}
defer listener.Close()

for {
    conn, err := listener.Accept()
    if err != nil {
        log.Println(err)
        continue
    }
    go handleConnection(conn)
}
```

---

# TCP Client Example

```go
conn, err := net.Dial("tcp", "localhost:8080")
if err != nil {
    log.Fatal(err)
}
defer conn.Close()

fmt.Fprintf(conn, "Hello Server!")

buffer := make([]byte, 1024)
n, err := conn.Read(buffer)
```

---

# UDP Server Example

```go
addr, err := net.ResolveUDPAddr("udp", ":8081")
if err != nil {
    log.Fatal(err)
}

conn, err := net.ListenUDP("udp", addr)
if err != nil {
    log.Fatal(err)
}
defer conn.Close()

buffer := make([]byte, 1024)
```

---

# UDP Client Example

```go
addr, err := net.ResolveUDPAddr("udp", "localhost:8081")
if err != nil {
    log.Fatal(err)
}

conn, err := net.DialUDP("udp", nil, addr)
if err != nil {
    log.Fatal(err)
}
defer conn.Close()
```

---

# Connection Handling

```go
func handleConnection(conn net.Conn) {
    defer conn.Close()
    
    // Set deadlines
    conn.SetDeadline(time.Now().Add(10 * time.Second))
    
    // Read data
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    
    // Process data
    // Send response
}
```

---

# Error Handling

```go
// Common network errors
type net.Error interface {
    error
    Timeout() bool   // Is it a timeout?
    Temporary() bool // Is it temporary?
}

if netErr, ok := err.(net.Error); ok {
    if netErr.Temporary() {
        // Retry operation
    }
}
```

---

# Best Practices

1. Always close connections
2. Set appropriate timeouts
3. Handle errors gracefully
4. Use goroutines for concurrency
5. Implement proper logging

---

# Security Considerations

1. TLS/SSL encryption
2. Input validation
3. Rate limiting
4. Access control
5. Error handling

---

# Exercise Time!

1. Create TCP chat server
2. Implement UDP data broadcast
3. Handle multiple clients
4. Add error recovery

---

# Questions?

Let's practice with hands-on examples!
