package postgres

import (
	"finance-backend/db"
	"fmt"

	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
)

type TagStore struct {
	*sqlx.DB
}

func NewTagStore(db *sqlx.DB) *TagStore {
	return &TagStore{
		DB: db,
	}
}

func (s *TagStore) Tag(id uuid.UUID) (db.Tag, error) {
	var entry db.Tag
	if err := s.Get(&entry, `SELECT * FROM tags WHERE id=$1`, id); err != nil {
		return db.Tag{}, fmt.Errorf("error getting tag: %w", err)
	}
}

func (s *TagStore) Tags() ([]db.Tag, error) {
	var tags []db.Tag
	if err := s.Select(&tags, `SELECT * FROM tags`); err != nil {
		return []db.Tag{}, fmt.Errorf("error getting tags: %w", err)
	}
	return tags, nil
}

func (s *TagStore) CreateTag(tag *db.Tag) error {
	if err := s.Get(tag, `INSERT INTO tags VALUES ($1, $2) RETURNING *`,
		tag.ID,
		tag.Name,
	); err != nil {
		return fmt.Errorf("error creating Tag: %w", err)
	}
	return nil
}

func (s *TagStore) UpdateTag(tag *db.Tag) error {
	if err := s.Get(tag, `UPDATE tags SET name=$2) WHERE id=$1`,
		tag.ID,
		tag.Name,
	); err != nil {
		return fmt.Errorf("error updating Tag: %w", err)
	}
	return nil
}

func (s *TagStore) DeleteTag(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM tags WHERE id=$1`, id); err != nil {
		return fmt.Errorf("error deleting Tag: %w", err)
	}
	return nil
}
