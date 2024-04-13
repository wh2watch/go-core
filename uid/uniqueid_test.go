package uid_test

import (
	"math/rand"
	"testing"
	"time"

	"core/uid"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
)

func TestNewUIDDoNotDuplicateIDs(t *testing.T) {
	uid1 := uid.NewUID()
	uid2 := uid.NewUID()
	assert.NotEqual(t, uid1, uid2)
}

func TestIsValidUID(t *testing.T) {
	tm := time.Unix(1000000, 0)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(tm.UnixNano())), 0)
	validULID := ulid.MustNew(ulid.Timestamp(tm), entropy).String()

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid ULID", validULID, true},
		{"Invalid ULID - Empty String", "", false},
		{"Invalid ULID - Random String", "not-a-valid-ulid", false},
		{"Invalid ULID - Partial ULID", validULID[:10], false},
		{"Invalid ULID - Overlong ULID", validULID + "123", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, uid.IsValidUID(tc.input))
		})
	}
}
