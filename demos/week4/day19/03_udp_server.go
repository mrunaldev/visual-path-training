package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Create a UDP listener
	addr, err := net.ResolveUDPAddr("udp", ":8081")
	if err != nil {
		log.Fatal("Error resolving address:", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer conn.Close()

	fmt.Println("UDP Server is listening on :8081")

	buffer := make([]byte, 1024)
	for {
		// Read incoming data
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("Error reading from UDP:", err)
			continue
		}

		message := string(buffer[:n])
		fmt.Printf("Received '%s' from %v\n", message, remoteAddr)

		// Echo back to client
		response := fmt.Sprintf("Server received: %s", message)
		_, err = conn.WriteToUDP([]byte(response), remoteAddr)
		if err != nil {
			log.Printf("Error sending response to %v: %v\n", remoteAddr, err)
		}
	}
}
