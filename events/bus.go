package events

import (
	"context"
)

type Handler[E DomainEvent] func(ctx context.Context, event E)

//go:generate mockery --name Bus --filename bus.go
type Bus interface {
	// Dispatch sends the given events to the appropriate event handlers.
	//
	// ctx: the context in which the events are dispatched.
	// events: variadic parameter representing the DomainEvents to be dispatched.
	Dispatch(ctx context.Context, events ...DomainEvent)
}
