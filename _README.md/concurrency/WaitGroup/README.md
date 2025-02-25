```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// ============= #### =============
	// ========== Concurrency =========
	// ============= #### =============

	// =========== Example 1 ==========
	// used to wait for multiple goroutines to finish
	// before the main function exits.
	var wg sync.WaitGroup

	wg.Add(1) // Step 1: wg = 1

	go func() {
		defer wg.Done() // Step 2: wg will decrement when this goroutine exits
		for i := 1; i < 200; i++ {
			fmt.Println("1 ====== go-routine: ", i)
			time.Sleep(time.Duration(rand.Intn(i)) * time.Millisecond)
		}
	}()

	wg.Add(1) // Step 3: wg = 2

	go func() {
		defer wg.Done() // Step 4: wg will decrement when this goroutine exits
		for i := 1; i < 200; i++ {
			fmt.Println("2 ###### go-routine: ", i)
			time.Sleep(time.Duration(rand.Intn(i)) * time.Millisecond)
		}
	}()

	// Step 5: Both goroutines are running concurrently

	wg.Wait() // Step 6: Waits until wg = 0

	fmt.Println("Done") // Step 7: Executed after both goroutines finish
}

```