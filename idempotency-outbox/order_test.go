package main

import (
	"context"
	"errors"
	"testing"
)

type mockUnitOfWork struct {
	repo *mockOrderRepo
}

func (m *mockUnitOfWork) Do(ctx context.Context, fn func(txRepo OrderRepository) error) error {

	snapshotOrders := make(map[string]Order)
	for k, v := range m.repo.orders {
		snapshotOrders[k] = v
	}
	snapshotOutbox := make([]OutboxMessage, len(m.repo.outbox))
	copy(snapshotOutbox, m.repo.outbox)

	err := fn(m.repo)
	if err != nil {
		m.repo.orders = snapshotOrders
		m.repo.outbox = snapshotOutbox
		return err
	}

	return nil
}

type mockOrderRepo struct {
	orders       map[string]Order
	outbox       []OutboxMessage
	failOnOutbox bool
}

func (m *mockOrderRepo) SaveOrder(ctx context.Context, o *Order) error {
	m.orders[o.ID] = *o
	return nil
}

func (m *mockOrderRepo) SaveOutbox(ctx context.Context, msg *OutboxMessage) error {
	if m.failOnOutbox {
		return ErrOutboxFailed
	}
	m.outbox = append(m.outbox, *msg)
	return nil
}

func (m *mockOrderRepo) FindOrderById(ctx context.Context, id string) (*Order, error) {
	o, exists := m.orders[id]
	if !exists {
		return nil, ErrOrderNotFound
	}
	return &o, nil
}

func (m *mockOrderRepo) FindByID(ctx context.Context, id string) (*Order, error) {
	return m.FindOrderById(ctx, id)
}

func TestCreateOrder_OutboxFailureRollsBackAggregate(t *testing.T) {
	ctx := context.Background()
	repo := &mockOrderRepo{
		orders:       make(map[string]Order),
		outbox:       make([]OutboxMessage, 0),
		failOnOutbox: true, // Force write failure on the event
	}
	uow := &mockUnitOfWork{repo: repo}
	useCase := NewCreateOrderUseCase(uow)

	cmd := CreateOrderCommand{
		OrderID:        "ord-1234",
		UserID:         "usr-9999",
		TotalCents:     15000,
		IdempotencyKey: "idem-key-abc",
	}

	err := useCase.Execute(ctx, cmd)

	if !errors.Is(err, ErrOutboxFailed) {
		t.Fatalf("expected ErrOutboxFailed, got: %v", err)
	}

	_, err = repo.FindByID(ctx, cmd.OrderID)
	if !errors.Is(err, ErrOrderNotFound) {
		t.Fatalf("expected order to be rolled back and not found, but query returned: %v", err)
	}

	if len(repo.outbox) != 0 {
		t.Fatalf("expected outbox to be empty, found %d records", len(repo.outbox))
	}
}
