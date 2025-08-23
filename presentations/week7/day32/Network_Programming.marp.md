---
marp: true
theme: default
paginate: true
---

# Advanced Network Programming in Go
## Day 32: Network Protocols and Implementation

---

# Overview

1. TCP/IP Programming
2. UDP Communication
3. HTTP/2 and gRPC
4. WebSocket Implementation
5. Network Security
6. Custom Protocol Design

---

# TCP/IP Programming

1. Basic TCP Server
```go
listener, err := net.Listen("tcp", ":8080")
for {
    conn, err := listener.Accept()
    go handleConnection(conn)
}
```

2. Connection Handling
```go
func handleConnection(conn net.Conn) {
    defer conn.Close()
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        // Process data
    }
}
```

---

# TCP Client Implementation

1. Connection
```go
conn, err := net.Dial("tcp", "localhost:8080")
```

2. Data Transfer
```go
type Message struct {
    Type    string `json:"type"`
    Payload []byte `json:"payload"`
}

encoder := json.NewEncoder(conn)
encoder.Encode(message)
```

---

# UDP Communication

1. UDP Server
```go
addr, _ := net.ResolveUDPAddr("udp", ":8081")
conn, _ := net.ListenUDP("udp", addr)

buffer := make([]byte, 1024)
for {
    n, remoteAddr, _ := conn.ReadFromUDP(buffer)
    go handleDatagram(buffer[:n], remoteAddr)
}
```

2. UDP Client
```go
addr, _ := net.ResolveUDPAddr("udp", "localhost:8081")
conn, _ := net.DialUDP("udp", nil, addr)
conn.Write([]byte("message"))
```

---

# HTTP/2 and gRPC

1. HTTP/2 Server
```go
server := &http.Server{
    Addr:    ":8443",
    Handler: handler,
    TLSConfig: &tls.Config{
        NextProtos: []string{"h2"},
    },
}
```

2. gRPC Service
```protobuf
service MyService {
    rpc StreamData (stream Request) 
        returns (stream Response);
}
```

---

# WebSocket Implementation

1. WebSocket Server
```go
upgrader := websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }
    handleWsConnection(conn)
}
```

2. WebSocket Client
```go
conn, _, err := websocket.DefaultDialer.Dial(
    "ws://localhost:8080/ws", nil)
```

---

# Network Security

1. TLS Configuration
```go
cert, _ := tls.LoadX509KeyPair("cert.pem", "key.pem")
config := &tls.Config{
    Certificates: []tls.Certificate{cert},
    MinVersion:  tls.VersionTLS12,
}
```

2. Certificate Pinning
```go
config.VerifyPeerCertificate = func(
    rawCerts [][]byte, 
    verifiedChains [][]*x509.Certificate) error {
    // Verify certificate fingerprint
}
```

---

# Custom Protocol Design

1. Protocol Definition
```go
type Protocol struct {
    Version    uint8
    Command    uint8
    PayloadLen uint16
    Payload    []byte
    Checksum   uint32
}
```

2. Frame Processing
```go
func (p *Protocol) Marshal() []byte {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.BigEndian, p.Version)
    // Write other fields
    return buf.Bytes()
}
```

---

# Error Handling

1. Network Errors
```go
type NetworkError struct {
    Op  string
    Err error
}

func (e *NetworkError) Error() string {
    return fmt.Sprintf("%s: %v", e.Op, e.Err)
}
```

2. Retry Logic
```go
func withRetry(fn func() error) error {
    for i := 0; i < maxRetries; i++ {
        if err := fn(); err == nil {
            return nil
        }
        time.Sleep(backoff(i))
    }
    return ErrMaxRetriesExceeded
}
```

---

# Connection Pool

1. Pool Implementation
```go
type Pool struct {
    conns chan net.Conn
    addr  string
    mu    sync.RWMutex
}

func (p *Pool) Get() net.Conn {
    select {
    case conn := <-p.conns:
        return conn
    default:
        return p.dial()
    }
}
```

---

# Performance Optimization

1. Zero-Copy
```go
func sendFile(conn net.Conn, file *os.File) error {
    _, err := syscall.Sendfile(
        int(conn.(*net.TCPConn).File()),
        int(file.Fd()),
        nil,
        0)
    return err
}
```

2. Buffer Pool
```go
pool := sync.Pool{
    New: func() interface{} {
        return make([]byte, 32*1024)
    },
}
```

---

# Monitoring and Metrics

1. Connection Metrics
```go
type Metrics struct {
    ActiveConns    int64
    BytesReceived  int64
    BytesSent     int64
    ErrorCount    int64
}
```

2. Prometheus Integration
```go
var (
    activeConnections = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Name: "active_connections",
            Help: "Number of active connections",
        })
)
```

---

# Best Practices

1. Always use timeouts
2. Implement proper error handling
3. Use connection pooling
4. Monitor connection metrics
5. Implement security measures
6. Handle graceful shutdown

---

# Hands-on Exercise

Build a network application with:
1. Custom protocol
2. Connection pooling
3. TLS security
4. Metrics collection
5. Error handling

---

# Key Takeaways

1. Network programming requires careful error handling
2. Security is critical
3. Performance optimization is important
4. Monitoring is essential
5. Protocol design matters

---

# Next Steps

1. Advanced Protocol Design
2. Network Security Patterns
3. Performance Tuning
4. Custom Transport Layers
5. Advanced Monitoring
