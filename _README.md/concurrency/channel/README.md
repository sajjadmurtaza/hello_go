```go
package main

import (
	"fmt"
)

func main() {
	// ============= #### =============
	// ========== Concurrency =========
	// ============= #### =============

	// =========== Example 2 ==========
	// channel
	//  allows goroutines to communicate with each other by sending
	// and receiving values.

	// Sending a value to the channel: ch <- value
	// Receiving a value from the channel: value := <- ch
	c := make(chan int)

	go func() {
		sum := 0

		for i := 1; i < 200; i++ {
			sum += i
			fmt.Println("1:=========")
		}

		c <- sum
	}()

	go func() {
		multi := 1

		for i := 1; i < 50; i++ {
			multi *= i
			fmt.Println("2----------")

		}

		c <- multi
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("Output:", <-c)
	}
}

```