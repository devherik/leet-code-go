package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type CreateOrderCommand struct {
	OrderID        string
	UserID         string
	TotalCents     int64
	IdempotencyKey string
}

type CreateOrderUseCase struct {
	uow UnitOfWork
}

func NewCreateOrderUseCase(uow UnitOfWork) *CreateOrderUseCase {
	return &CreateOrderUseCase{uow: uow}
}

func (uc *CreateOrderUseCase) Execute(ctx context.Context, cmd CreateOrderCommand) error {
	newOrder := &Order{
		ID:         cmd.OrderID,
		UserID:     cmd.UserID,
		TotalCents: cmd.TotalCents,
		Status:     "CREATED",
	}

	payload, err := json.Marshal(map[string]any{
		"order_id":    newOrder.ID,
		"user_id":     newOrder.UserID,
		"total_cents": newOrder.TotalCents,
		"status":      newOrder.Status,
	})
	if err != nil {
		return fmt.Errorf("failed to serialize event payload: %w", err)
	}

	event := &OutboxMessage{
		ID:             fmt.Sprintf("evt-%s", cmd.OrderID),
		IdempotencyKey: cmd.IdempotencyKey,
		AggregateType:  "Order",
		AggregateID:    newOrder.ID,
		Payload:        payload,
		CreatedAt:      time.Now().UTC(),
	}

	return uc.uow.Do(ctx, func(txRepo OrderRepository) error {
		if err := txRepo.SaveOrder(ctx, newOrder); err != nil {
			return fmt.Errorf("failed to save order: %w", err)
		}

		if err := txRepo.SaveOutbox(ctx, event); err != nil {
			return fmt.Errorf("%w: %v", ErrOutboxFailed, err)
		}

		return nil
	})

}
