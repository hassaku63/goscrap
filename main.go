package main

import (
	"fmt"
	"os"
)

func main() {
	// before wire
	// message := NewMessage()
	// greeter := NewGreeter(message)
	// event := NewEvent(greeter)

	// event.Start()

	// after wire
	greeter, err := InitializeEvent()
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	greeter.Start()
}
