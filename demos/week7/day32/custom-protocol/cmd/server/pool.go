package main

import (
	"net"
	"sync"
)

// ConnectionPool manages a pool of network connections
type ConnectionPool struct {
	mu      sync.RWMutex
	conns   chan net.Conn
	size    int
	factory func() (net.Conn, error)
}

// NewConnectionPool creates a new connection pool
func NewConnectionPool(size int) *ConnectionPool {
	return &ConnectionPool{
		conns: make(chan net.Conn, size),
		size:  size,
	}
}

// Get retrieves a connection from the pool
func (p *ConnectionPool) Get() net.Conn {
	select {
	case conn := <-p.conns:
		if conn == nil {
			return nil
		}
		return conn
	default:
		return nil
	}
}

// Put returns a connection to the pool
func (p *ConnectionPool) Put(conn net.Conn) {
	if conn == nil {
		return
	}

	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.conns == nil {
		conn.Close()
		return
	}

	select {
	case p.conns <- conn:
	default:
		conn.Close()
	}
}

// Close closes the connection pool
func (p *ConnectionPool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.conns == nil {
		return
	}

	close(p.conns)
	for conn := range p.conns {
		if conn != nil {
			conn.Close()
		}
	}
	p.conns = nil
}
