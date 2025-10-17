package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct{ DB *pgxpool.Pool }

func NewStore(db *pgxpool.Pool) *Store { return &Store{DB: db} }

func (s *Store) CreateOrder(ctx context.Context, o Order, items []OrderItem) error {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	_, err = tx.Exec(ctx, `INSERT INTO orders(id,user_id,status) VALUES($1,$2,$3)`, o.ID, o.UserID, o.Status)
	if err != nil {
		return err
	}
	for _, it := range items {
		_, err = tx.Exec(ctx, `INSERT INTO order_items(order_id,sku,quantity) VALUES($1,$2,$3)`, it.OrderID, it.SKU, it.Quantity)
		if err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}
