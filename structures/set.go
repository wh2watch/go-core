package structures

// Set is a set of comparable values
type Set[T comparable] map[T]struct{}

// NewSet creates a new set
func NewSet[T comparable]() Set[T] {
	return make(map[T]struct{})
}

// Add adds a value to the set
func (s Set[T]) Add(key T) {
	s[key] = struct{}{}
}

// Remove removes a value from the set
func (s Set[T]) Remove(key T) {
	delete(s, key)
}

// Contains returns true if the set contains the given value
func (s Set[T]) Contains(key T) bool {
	_, ok := s[key]
	return ok
}

// Size returns the number of elements in the set
func (s Set[T]) Size() int {
	return len(s)
}

// IsEmpty returns true if the set is empty
func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

// Values returns all the values in the set
func (s Set[T]) Values() []T {
	values := make([]T, 0, len(s))
	for key := range s {
		values = append(values, key)
	}
	return values
}
