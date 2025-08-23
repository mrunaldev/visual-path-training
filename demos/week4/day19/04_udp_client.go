package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Resolve UDP address
	addr, err := net.ResolveUDPAddr("udp", "localhost:8081")
	if err != nil {
		log.Fatal("Error resolving address:", err)
	}

	// Connect to server
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	fmt.Println("Connected to UDP server on localhost:8081")
	fmt.Println("Type your message and press Enter (type 'quit' to exit):")

	scanner := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, 1024)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		message := scanner.Text()
		if message == "quit" {
			break
		}

		// Send message
		_, err := conn.Write([]byte(message))
		if err != nil {
			log.Println("Error sending message:", err)
			continue
		}

		// Read response
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("Error reading response:", err)
			continue
		}

		fmt.Printf("Server: %s\n", string(buffer[:n]))
	}
}
