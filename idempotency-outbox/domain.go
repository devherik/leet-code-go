package main

import (
	"context"
	"errors"
	"time"
)

var (
	ErrOutboxFailed  = errors.New("failed to persist outbox event")
	ErrOrderNotFound = errors.New("order not found")
)

type OutboxMessage struct {
	ID             string
	IdempotencyKey string
	AggregateType  string
	AggregateID    string
	Payload        []byte
	CreatedAt      time.Time
}

type Order struct {
	ID         string
	UserID     string
	TotalCents int64
	Status     string
}

type UnitOfWork interface {
	Do(ctx context.Context, fn func(tRepo OrderRepository) error) error
}

type OrderRepository interface {
	SaveOrder(ctx context.Context, order *Order) error
	SaveOutbox(ctx context.Context, message *OutboxMessage) error
	FindOrderById(ctx context.Context, id string) (*Order, error)
}
