package uid

import "github.com/oklog/ulid/v2"

// UID represents an unique ID (ULID)
type UID struct {
	val ulid.ULID
}

// NewUID generates a new UID based on the current timestamp
func NewUID() UID {
	return UID{
		val: ulid.Make(),
	}
}

// String returns the string representation for the UID
func (i UID) String() string {
	return i.val.String()
}

// IsValidUID checks if the given string is a valid ULID
func IsValidUID(s string) bool {
	_, err := ulid.ParseStrict(s)
	return err == nil
}
