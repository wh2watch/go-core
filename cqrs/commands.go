package cqrs

import (
	"context"
)

// Command is a generic interface for commands
type Command any

// CommandHandler is a generic interface for command handlers
type CommandHandler[C Command] func(ctx context.Context, cmd C) error

// CommandBus is a generic interface for command bus
//
//go:generate mockery --name CommandBus --filename commandbus.go
type CommandBus interface {
	// Dispatch is a function that executes the given command.
	// It takes a context and a Command interface as parameters.
	Dispatch(ctx context.Context, command Command) error
}
