package order

import (
	"database/sql"

	"github.com/CuongDepay/go-ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	query := "INSERT INTO orders (userId, total, status, address) VALUES (?, ?, ?, ?)"
	res, err := s.db.Exec(query, order.UserID, order.Total, order.Status, order.Address)
	if err != nil {
		return 0, err
	}

	orderID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(orderID), nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	query := "INSERT INTO order_items (orderId, productId, quantity, price) VALUES (?, ?, ?, ?)"
	_, err := s.db.Exec(query, orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	if err != nil {
		return err
	}

	return nil
}
