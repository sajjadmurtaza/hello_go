// ============= ####### =============
// ============== main ==============
// ============= ####### =============

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// ============= ####### =============
// ============== funcs =============
// ============= ####### =============

// connectToServer simulates a streaming service connection
// It demonstrates using context for timeout control
func connectToServer(done chan bool) {
	// ============= ####### =============
	// ======== connection attempt =======
	// ============= ####### =============

	fmt.Println("Attempting connection...")

	// Simulate random connection success/failure
	if rand.Intn(100) > 50 {
		fmt.Println("OOPS! Connection failed")
		// Simulate hanging connection
		time.Sleep(1000 * time.Hour)
	} else {
		// Simulate successful connection delay
		time.Sleep(2 * time.Second)
		fmt.Println("Connection successful")
	}

	// Signal completion through channel
	done <- true
}

// ============= ####### =============
// ============== main ==============
// ============= ####### =============

func main() {
	// ============= ####### =============
	// ======== context timeout =========
	// ============= ####### =============

	// Create context with 3s timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Ensure context is cancelled when done

	// Create channel for connection status
	done := make(chan bool)

	// ============= ####### =============
	// ======== connection logic ========
	// ============= ####### =============

	// Start connection attempt in goroutine
	go connectToServer(done)

	// Wait for either success or timeout
	select {
	case <-done:
		fmt.Println("Connection Successful")
	case <-ctx.Done():
		fmt.Println("Connection attempt timed out")
	}
}
