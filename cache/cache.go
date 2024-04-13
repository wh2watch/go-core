package cache

import (
	"context"
	"time"
)

type Cache[T any] interface {
	// Get returns the value associated with the given key
	Get(context context.Context, key string) (T, error)
	// Set associates the given value with the given key
	Set(context context.Context, key string, value T, opts ...SetOption) error
}

// SetOptions is a struct that contains options for a Set operation
type SetOptions struct {
	TTL time.Duration
}

// SetOption is a function that sets options for a Set operation
type SetOption func(options *SetOptions)

// WithTTL returns a SetOption function that sets the time to live (TTL) for a key
func WithTTL(ttl time.Duration) SetOption {
	return func(options *SetOptions) {
		options.TTL = ttl
	}
}

// CollectSetOptions returns a new SetOptions with the given options
func CollectSetOptions(options ...SetOption) SetOptions {
	opts := SetOptions{}
	for _, option := range options {
		option(&opts)
	}
	return opts
}
