package postgres

import (
	"finance-backend/db"
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq" // To register the driver.
)

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Store{
		TagStore:      NewTagStore(db),
		CategoryStore: NewCategoryStore(db),
		EntryStore:    NewEntryStore(db),
		UserStore:     NewUserStore(db),
	}, nil
}

type Store struct {
	db.TagStore
	db.EntryStore
	db.UserStore
	db.CategoryStore
}
