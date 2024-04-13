package cqrs

import "context"

// Query is a generic interface for queries
type Query any

// Response is a generic interface for query responses
type Response any

// QueryHandler is a generic interface for query handlers
type QueryHandler[Q Query, R Response] func(ctx context.Context, qry Q) (R, error)

// QueryBus is a generic interface for query bus
//
//go:generate mockery --name QueryBus --filename querybus.go
type QueryBus interface {
	// Dispatch method sends the query and wait for the response which is saved in the res param. res param must be a pointer.
	Dispatch(ctx context.Context, qry Query, res Response) error
}
