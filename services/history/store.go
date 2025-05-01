package history

import (
	"database/sql"
	"fmt"

	"github.com/aarav345/ecom-go/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateHistory(history types.History) error {
	fmt.Println(history)
	return nil
}
