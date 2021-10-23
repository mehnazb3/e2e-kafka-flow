package main

import (
	"fmt"

	"github.com/mehnazb3/e2e-kafka-flow/producer"
)

func main() {
	fmt.Println("Starting!!!")
	fmt.Println("Initializing Producer")
	producer.Init()
	// Config
	// Send
	// Get message partition and offset details
	// Logging
	fmt.Println("Initializing Consumer")
	// Config
	// Subscribe to topics
	// Read from topics
	// Logging
	// Sync and Async
	fmt.Println("Exiting!!!")
}
