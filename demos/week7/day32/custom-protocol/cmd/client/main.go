package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"custom-protocol/protocol"
)

// Client represents a protocol client
type Client struct {
	conn net.Conn
}

// NewClient creates a new client
func NewClient(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}

	return &Client{conn: conn}, nil
}

// Close closes the client connection
func (c *Client) Close() error {
	return c.conn.Close()
}

// SendPing sends a ping message
func (c *Client) SendPing() error {
	frame, err := protocol.NewFrame(protocol.CmdPing, []byte("ping"))
	if err != nil {
		return err
	}

	if err := c.sendFrame(frame); err != nil {
		return err
	}

	response, err := protocol.Unmarshal(c.conn)
	if err != nil {
		return err
	}

	log.Printf("Received pong: %s", string(response.Payload))
	return nil
}

// SendMessage sends a message
func (c *Client) SendMessage(msg string) error {
	frame, err := protocol.NewFrame(protocol.CmdMessage, []byte(msg))
	if err != nil {
		return err
	}

	if err := c.sendFrame(frame); err != nil {
		return err
	}

	response, err := protocol.Unmarshal(c.conn)
	if err != nil {
		return err
	}

	log.Printf("Received echo: %s", string(response.Payload))
	return nil
}

// SendFile sends a file
func (c *Client) SendFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	frame, err := protocol.NewFrame(protocol.CmdFile, data)
	if err != nil {
		return err
	}

	if err := c.sendFrame(frame); err != nil {
		return err
	}

	response, err := protocol.Unmarshal(c.conn)
	if err != nil {
		return err
	}

	log.Printf("File transfer response: %s", string(response.Payload))
	return nil
}

// sendFrame sends a frame to the server
func (c *Client) sendFrame(frame *protocol.Frame) error {
	data, err := frame.Marshal()
	if err != nil {
		return err
	}

	_, err = c.conn.Write(data)
	return err
}

func main() {
	addr := flag.String("addr", "localhost:8080", "Server address")
	message := flag.String("message", "", "Message to send")
	file := flag.String("file", "", "File to send")
	flag.Parse()

	client, err := NewClient(*addr)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Send ping every second
	go func() {
		for {
			if err := client.SendPing(); err != nil {
				log.Printf("Ping failed: %v", err)
				return
			}
			time.Sleep(time.Second)
		}
	}()

	// Send message if provided
	if *message != "" {
		if err := client.SendMessage(*message); err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	}

	// Send file if provided
	if *file != "" {
		if err := client.SendFile(*file); err != nil {
			log.Printf("Failed to send file: %v", err)
		}
	}

	// Keep the main goroutine running
	select {}
}
