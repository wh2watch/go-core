package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Event struct {
	EventBase
}

func TestEventBase_SetTime(t *testing.T) {
	evt := Event{}
	now := time.Now()
	evt.SetTime(now)
	assert.Equal(t, now, evt.OccurredAt())
}
