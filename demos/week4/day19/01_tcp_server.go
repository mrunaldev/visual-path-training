package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Create a TCP listener
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()

	fmt.Println("Server is listening on :8080")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// Handle each connection in a goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to store received data
	buffer := make([]byte, 1024)

	for {
		// Read incoming data
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error reading from connection:", err)
			return
		}

		// Echo received data back to client
		message := string(buffer[:n])
		fmt.Printf("Received: %s\n", message)

		// Send response
		response := fmt.Sprintf("Server received: %s", message)
		_, err = conn.Write([]byte(response))
		if err != nil {
			log.Println("Error writing to connection:", err)
			return
		}
	}
}
