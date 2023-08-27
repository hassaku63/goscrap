//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewMessage)
	return Event{}, nil
}
