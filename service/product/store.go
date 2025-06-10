package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/CuongDepay/go-ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]*types.Product, error) {
	query := "SELECT * FROM products"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]*types.Product, 0)
	for rows.Next() {
		product, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (s *Store) CreateProduct(product types.CreateProductPayload) error {
	query := "INSERT INTO products (name, description, image, price, quantity) VALUES (?, ?, ?, ?, ?)"
	_, err := s.db.Exec(query, product.Name, product.Description, product.Image, product.Price, product.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateProduct(product types.Product) error {
	query := "UPDATE products SET name = ?, description = ?, image = ?, price = ?, quantity = ? WHERE id = ?"
	_, err := s.db.Exec(query, product.Name, product.Description, product.Image, product.Price, product.Quantity, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetProductByID(productID int) (*types.Product, error) {
	query := "SELECT * FROM products WHERE id = ?"
	rows, err := s.db.Query(query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	p := new(types.Product)
	for rows.Next() {
		p, err = scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (s *Store) GetProductsByIDs(productIDs []int) ([]types.Product, error) {
	placeholders := strings.Repeat("?,", len(productIDs)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeholders)

	args := make([]interface{}, len(productIDs))
	for i, id := range productIDs {
		args[i] = id
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]types.Product, 0)
	for rows.Next() {
		product, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}

	return products, nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Image, &product.Price, &product.Quantity, &product.CreatedAt)
	if err != nil {
		return nil, err
	}

	return product, nil
}
