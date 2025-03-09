package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func fetchData(ctx context.Context, userID int) (bool, error) {
	time.Sleep(400 * time.Millisecond)

	if ctx.Err() == context.DeadlineExceeded {
		return false, fmt.Errorf("fetchData: context deadline exceeded")
	}

	return true, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)

	defer cancel()

	userID := 1

	_, err := fetchData(ctx, userID)
	if err != nil {
		log.Fatalf("error fetching data: %v", err)
	}

	fmt.Printf("success with UserID: %d\n", userID)
}
