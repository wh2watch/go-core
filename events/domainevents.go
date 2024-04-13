package events

import "time"

// DomainEvent is the base interface for all domain events.
type DomainEvent interface {
	// OccurredAt returns the time when the event was created.
	OccurredAt() time.Time
}

// timeSetable is an interface that allows to set the time when the event was created.
// It is used by the AggregateRootBase to set the time when an event is registered
type timeSetable interface {
	// SetTime sets the time when the event was created.
	SetTime(time.Time)
}

// EventBase is the base struct for all domain events.
// It implements both DomainEvent and timeSetable interfaces
type EventBase struct {
	occurredAt time.Time
}

// OccurredAt returns the time when the event was created.
func (e EventBase) OccurredAt() time.Time {
	return e.occurredAt
}

// SetTime sets the time when the event was created.
func (e *EventBase) SetTime(t time.Time) {
	e.occurredAt = t
}
