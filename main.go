package main

import (
	"fmt"
	"os"

	"github.com/hassaku63/goscrap/greet"
)

func main() {
	greeter, err := greet.InitializeEvent("Hi there!")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	greeter.Start()
}
