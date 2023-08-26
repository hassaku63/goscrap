package main

func main() {
	// before wire
	// message := NewMessage()
	// greeter := NewGreeter(message)
	// event := NewEvent(greeter)

	// event.Start()

	// after wire
	e := InitializeEvent()
	e.Start()
}
