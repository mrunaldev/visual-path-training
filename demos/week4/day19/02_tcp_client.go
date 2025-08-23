package tcpclient

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Connect to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	fmt.Println("Connected to server on localhost:8080")
	fmt.Println("Type your message and press Enter (type 'quit' to exit):")

	// Create scanner for reading user input
	scanner := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, 1024)

	for {
		// Read user input
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		message := scanner.Text()
		if message == "quit" {
			break
		}

		// Send message to server
		_, err := fmt.Fprintf(conn, message)
		if err != nil {
			log.Println("Error sending message:", err)
			break
		}

		// Read server response
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error reading response:", err)
			break
		}

		// Print server response
		fmt.Printf("Server: %s\n", string(buffer[:n]))
	}
}
