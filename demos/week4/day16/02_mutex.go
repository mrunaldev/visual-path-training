// Package sync_examples demonstrates Mutex usage
package sync_examples

import (
	"fmt"
	"sync"
)

// Counter represents a thread-safe counter
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment adds 1 to the counter safely
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current count safely
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// Bank represents a simple bank account
type Bank struct {
	mu      sync.Mutex
	balance int
}

// Deposit adds money to the account
func (b *Bank) Deposit(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.balance += amount
	fmt.Printf("Deposited %d, new balance: %d\n", amount, b.balance)
}

// Withdraw removes money from the account if sufficient funds exist
func (b *Bank) Withdraw(amount int) bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.balance >= amount {
		b.balance -= amount
		fmt.Printf("Withdrew %d, new balance: %d\n", amount, b.balance)
		return true
	}

	fmt.Printf("Insufficient funds for withdrawal of %d (balance: %d)\n",
		amount, b.balance)
	return false
}

// GetBalance returns the current balance
func (b *Bank) GetBalance() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.balance
}

// Example usage in main package:
/*
func main() {
	// Counter example
	counter := &Counter{}
	var wg sync.WaitGroup

	// Launch multiple goroutines to increment
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Final count: %d\n", counter.Value())

	// Bank example
	bank := &Bank{balance: 100}

	// Simulate concurrent transactions
	wg.Add(4)
	go func() {
		defer wg.Done()
		bank.Deposit(50)
	}()
	go func() {
		defer wg.Done()
		bank.Withdraw(70)
	}()
	go func() {
		defer wg.Done()
		bank.Deposit(30)
	}()
	go func() {
		defer wg.Done()
		bank.Withdraw(150)
	}()

	wg.Wait()
	fmt.Printf("Final balance: %d\n", bank.GetBalance())
}
*/
