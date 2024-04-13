package events

import "time"

// AggregateRootBase is a trait supporting events handling to compose in any AggregateRoot
type AggregateRootBase struct {
	domainEvents []DomainEvent
}

// RegisterEvent registers a domain event and sets the event occurrence time if it's unset.
func (a *AggregateRootBase) RegisterEvent(evt DomainEvent) {
	if se, isSetable := evt.(timeSetable); isSetable && evt.OccurredAt().IsZero() {
		se.SetTime(time.Now())
	}
	a.domainEvents = append(a.domainEvents, evt)
}

// FlushEvents returns and clears the domain events stored in the aggregate.
func (a *AggregateRootBase) FlushEvents() []DomainEvent {
	events := a.domainEvents
	a.domainEvents = nil
	return events
}
