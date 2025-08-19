package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Quote represents a programming quote with its author
type Quote struct {
	Text   string
	Author string
}

// getQuotes returns a slice of programming quotes
func getQuotes() []Quote {
	return []Quote{
		{
			Text:   "Don't comment bad code - rewrite it.",
			Author: "Brian Kernighan",
		},
		{
			Text:   "Simplicity is prerequisite for reliability.",
			Author: "Edsger W. Dijkstra",
		},
		{
			Text:   "The best way to predict the future is to invent it.",
			Author: "Alan Kay",
		},
		{
			Text:   "Talk is cheap. Show me the code.",
			Author: "Linus Torvalds",
		},
		{
			Text:   "Programming isn't about what you know; it's about what you can figure out.",
			Author: "Chris Pine",
		},
	}
}

// getRandomQuote returns a random quote from our collection
func getRandomQuote(quotes []Quote) Quote {
	randomIndex := rand.Intn(len(quotes))
	return quotes[randomIndex]
}

// printQuote displays a quote in a formatted way
func printQuote(quote Quote) {
	fmt.Println("╔════════════════════════════════════════════")
	fmt.Printf("║ %s\n", quote.Text)
	fmt.Printf("║ - %s\n", quote.Author)
	fmt.Println("╚════════════════════════════════════════════")
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Welcome message
	fmt.Println("Welcome to Quote of the Day!")
	fmt.Println("----------------------------")

	// Get all quotes
	quotes := getQuotes()

	// Get and display a random quote
	todaysQuote := getRandomQuote(quotes)
	printQuote(todaysQuote)

	// Interactive element
	var input string
	fmt.Print("\nWould you like another quote? (yes/no): ")
	fmt.Scan(&input)

	if input == "yes" {
		newQuote := getRandomQuote(quotes)
		// Make sure we don't get the same quote twice
		for newQuote == todaysQuote {
			newQuote = getRandomQuote(quotes)
		}
		printQuote(newQuote)
	}

	fmt.Println("\nThank you for trying our Go demo!")
}
