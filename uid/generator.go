package uid

//go:generate mockery --name Generator --filename generator.go
type Generator interface {
	// Generate returns a string.
	Generate() string
}

// UIDGenerator is an implementation of the Generator interface.
type UIDGenerator struct{}

// NewUIDGenerator creates a new instance of UIDGenerator.
func NewUIDGenerator() UIDGenerator {
	return UIDGenerator{}
}

// Generate generates a new unique identifier string.
func (g UIDGenerator) Generate() string {
	return NewUID().String()
}
