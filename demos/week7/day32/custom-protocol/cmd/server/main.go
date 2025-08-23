package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"custom-protocol/protocol"

	"github.com/prometheus/client_golang/prometheus"
)

// Metrics
var (
	activeConnections = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "active_connections",
		Help: "Number of active connections",
	})

	totalMessages = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "total_messages",
		Help: "Total number of messages processed",
	})

	messageLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "message_latency_seconds",
		Help:    "Message processing latency in seconds",
		Buckets: prometheus.DefBuckets,
	})
)

// Server represents the protocol server
type Server struct {
	listener net.Listener
	pool     *ConnectionPool
	done     chan struct{}
	wg       sync.WaitGroup
}

// NewServer creates a new server instance
func NewServer(addr string, poolSize int) (*Server, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create listener: %v", err)
	}

	pool := NewConnectionPool(poolSize)

	return &Server{
		listener: listener,
		pool:     pool,
		done:     make(chan struct{}),
	}, nil
}

// Start starts the server
func (s *Server) Start() error {
	defer s.listener.Close()

	// Register metrics
	prometheus.MustRegister(activeConnections)
	prometheus.MustRegister(totalMessages)
	prometheus.MustRegister(messageLatency)

	log.Printf("Server started on %s", s.listener.Addr())

	for {
		select {
		case <-s.done:
			return nil
		default:
			conn, err := s.listener.Accept()
			if err != nil {
				log.Printf("Failed to accept connection: %v", err)
				continue
			}

			s.wg.Add(1)
			go s.handleConnection(conn)
		}
	}
}

// Stop stops the server gracefully
func (s *Server) Stop() {
	close(s.done)
	s.wg.Wait()
	s.pool.Close()
}

// handleConnection handles a client connection
func (s *Server) handleConnection(conn net.Conn) {
	defer s.wg.Done()
	defer conn.Close()

	// Update metrics
	activeConnections.Inc()
	defer activeConnections.Dec()

	log.Printf("New connection from %s", conn.RemoteAddr())

	for {
		select {
		case <-s.done:
			return
		default:
			// Set read deadline
			conn.SetReadDeadline(time.Now().Add(30 * time.Second))

			// Read frame
			frame, err := protocol.Unmarshal(conn)
			if err != nil {
				log.Printf("Failed to read frame: %v", err)
				return
			}

			start := time.Now()

			// Process frame
			response, err := s.processFrame(frame)
			if err != nil {
				log.Printf("Failed to process frame: %v", err)
				return
			}

			// Update metrics
			totalMessages.Inc()
			messageLatency.Observe(time.Since(start).Seconds())

			// Send response
			responseBytes, err := response.Marshal()
			if err != nil {
				log.Printf("Failed to marshal response: %v", err)
				return
			}

			if _, err := conn.Write(responseBytes); err != nil {
				log.Printf("Failed to send response: %v", err)
				return
			}
		}
	}
}

// processFrame processes a protocol frame
func (s *Server) processFrame(frame *protocol.Frame) (*protocol.Frame, error) {
	switch frame.Command {
	case protocol.CmdPing:
		return protocol.NewFrame(protocol.CmdPing, []byte("pong"))

	case protocol.CmdMessage:
		// Echo the message back
		return protocol.NewFrame(protocol.CmdMessage, frame.Payload)

	case protocol.CmdFile:
		// Store file in connection pool
		conn := s.pool.Get()
		if conn != nil {
			defer s.pool.Put(conn)
			if _, err := conn.Write(frame.Payload); err != nil {
				return nil, err
			}
		}
		return protocol.NewFrame(protocol.CmdFile, []byte("file received"))

	default:
		return nil, protocol.ErrInvalidCommand
	}
}

func main() {
	addr := flag.String("addr", ":8080", "Server address")
	poolSize := flag.Int("pool-size", 10, "Connection pool size")
	flag.Parse()

	server, err := NewServer(*addr, *poolSize)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
