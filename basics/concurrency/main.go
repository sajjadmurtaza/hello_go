package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go count("Cat", ch)
	go count("Dog", ch)
	go count("Sheep", ch)

	// receive message from channel until it's closed
	for message := range ch {
		fmt.Println(message)
	}
}

func count(thing string, ch chan string) {
	for i := 1; i <= 5; i++ {

		// send thing to channel
		ch <- thing

		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}

	close(ch)
}
