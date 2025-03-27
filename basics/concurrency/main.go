// ============= ####### =============
// ============== main ==============
// ============= ####### =============

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ============= ####### =============
// ============== funcs =============
// ============= ####### =============

// count demonstrates goroutine and channel usage
// It counts from 1 to 5 and sends the count to the channel
func count(thing string, ch chan string) {
	// ============= ####### =============
	// ========= send to channel =========
	// ============= ####### =============

	for i := 1; i <= 5; i++ {
		// Send the current count to the channel
		ch <- fmt.Sprintf("%s: %d", thing, i)

		// Random delay between 0-500ms
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}

	// Close the channel when done
	close(ch)
}

// ============= ####### =============
// ============== main ==============
// ============= ####### =============

func main() {
	// ============= ####### =============
	// ========= channel setup ==========
	// ============= ####### =============

	// Create a buffered channel with capacity 1
	ch := make(chan string, 1)

	// ============= ####### =============
	// ======== start goroutines ========
	// ============= ####### =============

	// Start three goroutines to count different things
	go count("Cat", ch)
	go count("Dog", ch)
	go count("Sheep", ch)

	// ============= ####### =============
	// ======= channel iteration ========
	// ============= ####### =============

	// Receive messages from channel until it's closed
	// This demonstrates channel iteration
	for message := range ch {
		fmt.Println(message)
	}
}
