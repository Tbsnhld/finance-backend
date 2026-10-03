package postgres

import (
	"finance-backend/db"
	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
)

type CategoryStore struct {
	*sqlx.DB
}

func NewCategoryStore(db *sqlx.DB) *CategoryStore {
	return &CategoryStore{
		DB: db,
	}
}

func (s *CategoryStore) Category(id uuid.UUID) (db.Category, error) {
	var category db.Category
	if err := s.Get(&category, `SELECT * FROM categories WHERE id=$1`, id); err != nil {
		return db.Category{}, fmt.Errorf("error getting category: %w", err)
	}
	return category, nil
}

func (s *CategoryStore) Categories() ([]db.Category, error) {
	var categorys []db.Category
	if err := s.Select(&categorys, `SELECT * FROM categories`); err != nil {
		return []db.Category{}, fmt.Errorf("error getting categories: %w", err)
	}
	return categorys, nil
}

func (s *CategoryStore) CreateCategory(category *db.Category) error {
	if err := s.Get(category, `INSERT INTO categories VALUES ($1, $2) RETURNING *`,
		category.ID,
		category.Name,
	); err != nil {
		return fmt.Errorf("error creating category: %w", err)
	}
	return nil
}

func (s *CategoryStore) UpdateCategory(category *db.Category) error {
	if err := s.Get(category, `UPDATE categories SET name=$2) WHERE id=$1`,
		category.ID,
		category.Name,
	); err != nil {
		return fmt.Errorf("error updating category: %w", err)
	}
	return nil
}

func (s *CategoryStore) DeleteCategory(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM categories WHERE id=$1`, id); err != nil {
		return fmt.Errorf("error deleting category: %w", err)
	}
	return nil
}
