package main

import (
	"context"
	"database/sql"
	"fmt"
)

type SQLOrderRepository struct {
	db *sql.DB
}

func NewSQLOrderRepository(db *sql.DB) *SQLOrderRepository {
	return &SQLOrderRepository{db: db}
}

func (r *SQLOrderRepository) CreateOrderWithEvent(
	ctx context.Context,
	order *Order,
	event *OutboxMessage,
) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return fmt.Errorf("failed to begin tx: %w", err)
	}

	defer tx.Rollback()

	const insertOrderQuery = `
		INSERT INTO orders (id, user_id, total_cents, status)
		VALUE ($1, $2, $3, $4)
		`

	_, err = tx.ExecContext(ctx, insertOrderQuery, order.ID, order.UserID, order.TotalCents, order.Status)
	if err != nil {
		return fmt.Errorf("failed to insert order: %w", err)
	}

	const insertOutboxQuery = `
		INSERT INTO outbox (id, idempotency_key, aggregated_type, aggregated_id, payload, created_at)
		VALUE ($1, $2, $3, $4, $5, $6)
		`

	_, err = tx.ExecContext(
		ctx,
		insertOutboxQuery,
		event.ID,
		event.IdempotencyKey,
		event.AggregateType,
		event.AggregateID,
		event.Payload,
		event.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to insert outbox event: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	return nil
}
