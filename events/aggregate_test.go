package events_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"core/events"
)

type DomainEvent1 struct {
	*events.EventBase
}

type DomainEvent2 struct {
	*events.EventBase
}

func TestAggregate(t *testing.T) {
	evt1 := DomainEvent1{&events.EventBase{}}
	evt2 := DomainEvent2{&events.EventBase{}}

	t.Run("Given an aggregate", func(t *testing.T) {
		sut := events.AggregateRootBase{}
		t.Run("When it register a succession of domain events", func(t *testing.T) {
			start := time.Now()
			sut.RegisterEvent(evt1)
			sut.RegisterEvent(evt2)
			end := time.Now()
			t.Run("Then FlushEvents returns the registered events", func(t *testing.T) {
				events := sut.FlushEvents()
				assert.Len(t, events, 2)
				assert.Equal(t, evt1, events[0])
				assert.Equal(t, evt2, events[1])
				assert.WithinRange(t, evt1.OccurredAt(), start, end)
				assert.WithinRange(t, evt2.OccurredAt(), start, end)
				t.Run("And a second call to FlushEvents returns empty", func(t *testing.T) {
					assert.Empty(t, sut.FlushEvents())
				})
			})
		})
	})
}
