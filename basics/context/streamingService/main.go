package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func connectToServer(done chan bool) {
	fmt.Println("Attempting connection...")

	if rand.Intn(100) > 50 {
		fmt.Println("OOPS! Connection failed")
		time.Sleep(1000 * time.Hour)
	} else {
		time.Sleep(2 * time.Second)
		fmt.Println("Connection successful")
	}

	done <- true
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	done := make(chan bool)
	go connectToServer(done)

	select {
	case <-done:
		fmt.Println("Connection Successful")
	case <-ctx.Done():
		fmt.Println("Connection attempt timed out")
	}
}
