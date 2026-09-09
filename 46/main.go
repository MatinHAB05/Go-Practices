package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// 1. Create an errgroup along with a derived Context
	g, ctx := errgroup.WithContext(context.Background())

	// Task 1: Fetch profile data
	g.Go(func() error {
		fmt.Println("Fetching user profile...")
		time.Sleep(500 * time.Millisecond)
		return nil // Success
	})

	// Task 2: Fetch financial data (simulating a failure)
	g.Go(func() error {
		fmt.Println("Fetching financial data...")
		time.Sleep(200 * time.Millisecond)
		return errors.New("financial database unavailable") // ❌ Error occurs
	})

	// Task 3: Fetch orders data
	g.Go(func() error {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Orders fetched successfully.")
			return nil
		case <-ctx.Done(): // 💡 Triggered as soon as another task returns an error!
			fmt.Println("Fetching orders canceled because another task failed.")
			return ctx.Err()
		}
	})

	// 2. Wait for all tasks to complete or for the first error to occur
	if err := g.Wait(); err != nil {
		fmt.Printf("❌ Group execution failed: %v\n", err)
	} else {
		fmt.Println("✅ All tasks completed successfully.")
	}
}
