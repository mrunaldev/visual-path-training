// Package main demonstrates importing and using external packages
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Configuration represents application settings
type Configuration struct {
	ServerName string   `json:"serverName"`
	Port       int      `json:"port"`
	Features   []string `json:"features"`
}

// String returns a formatted string representation
func (c Configuration) String() string {
	return fmt.Sprintf("Server: %s, Port: %d, Features: %s",
		c.ServerName,
		c.Port,
		strings.Join(c.Features, ", "))
}

func main() {
	// JSON configuration example
	jsonData := `{
		"serverName": "prod-server",
		"port": 8080,
		"features": ["logging", "metrics", "auth"]
	}`

	var config Configuration
	if err := json.Unmarshal([]byte(jsonData), &config); err != nil {
		log.Fatal("Error parsing configuration:", err)
	}

	fmt.Println("Configuration:", config)

	// Demonstrate string manipulation
	fmt.Println("Server name uppercase:", strings.ToUpper(config.ServerName))
	fmt.Println("Has logging feature:",
		strings.Contains(strings.Join(config.Features, ","), "logging"))
}
