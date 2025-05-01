package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/aarav345/ecom-go/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetProducts() ([]types.ProductWithInventory, error) {
	rows, err := s.db.Query(`
	SELECT p.id, p.name, p.description, p.image, p.price,
	COALESCE(i.quantity, 0) 
	FROM products as p 
	LEFT JOIN inventory as i 
	ON p.id = i.product_id`)

	if err != nil {
		return nil, err
	}

	var products []types.ProductWithInventory
	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) GetProductsByID(productIDs []int) ([]types.ProductWithInventory, error) {
	placeholders := strings.Repeat(",?", len(productIDs)-1)

	query := fmt.Sprintf("SELECT p.id, p.name, p.description, p.image, p.price, p.created_at, i.quantity FROM products as p LEFT JOIN inventory as i on p.id = i.product_id WHERE id IN (?%s)", placeholders)

	args := make([]interface{}, len(productIDs))
	for i, v := range productIDs {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := []types.ProductWithInventory{}
	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) UpdateProduct(product types.ProductWithInventory, updateProductFields bool) error {

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	if updateProductFields {
		_, err := tx.Exec("UPDATE products SET name = ?, price = ?, image = ?, description = ? WHERE id = ?", product.Name, product.Price, product.Image, product.Description, product.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	_, err = tx.Exec("UPDATE inventory SET quantity = ? WHERE product_id = ?", product.Quantity, product.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func scanRowIntoProduct(rows *sql.Rows) (*types.ProductWithInventory, error) {
	product := new(types.ProductWithInventory)

	if err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.CreatedAt,
		&product.Quantity,
	); err != nil {
		return nil, err
	}

	return product, nil
}
