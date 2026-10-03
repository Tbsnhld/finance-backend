package postgres

import (
	"finance-backend/db"
	"fmt"

	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
)

type EntryStore struct {
	*sqlx.DB
}

func NewEntryStore(db *sqlx.DB) *EntryStore {
	return &EntryStore{
		DB: db,
	}
}

func (s *EntryStore) Entry(id uuid.UUID) (db.Entry, error) {
	var entry db.Entry
	if err := s.Get(&entry, `SELECT * FROM entries WHERE id=$1`, id); err != nil {
		return db.Entry{}, fmt.Errorf("error getting entry: %w", err)
	}
}

func (s *EntryStore) Entries() ([]db.Entry, error) {
	var entries []db.Entry
	if err := s.Select(&entries, `SELECT * FROM entries`); err != nil {
		return []db.Entry{}, fmt.Errorf("error getting entries: %w", err)
	}
}

func (s *EntryStore) EntriesByUser(userId uuid.UUID) ([]db.Entry, error) {
	var entries_by_user []db.Entry
	if err := s.Select(&entries_by_user, `SELECT * FROM entries WHERE createdBy = $1`, userId); err != nil {
		return []db.Entry{}, fmt.Errorf("error getting entries for user $1: $2", userId, err)
	}
	return entries_by_user, nil
}

func (s *EntryStore) EntriesByCategory(categoryId uuid.UUID) ([]db.Entry, error) {
	var entries_by_category []db.Entry
	if err := s.Select(&entries_by_category, `SELECT * FROM entries WHERE categoryId = $1`, categoryId); err != nil {
		return []db.Entry{}, fmt.Errorf("error getting entries for user $1: $2", categoryId, err)
	}
	return entries_by_category, nil

}

func (s *EntryStore) EntriesByTag(tagId uuid.UUID) ([]db.Entry, error) {
	var entries_by_tag []db.Entry
	if err := s.Select(&entries_by_tag, `SELECT * FROM entries WHERE categoryId = $1`, tagId); err != nil {
		return []db.Entry{}, fmt.Errorf("error getting entries for user $1: $2", tagId, err)
	}
	return entries_by_tag, nil

}

func (s *EntryStore) CreateEntry(entry *db.Entry) error {
	if err := s.Get(entry, `INSERT INTO entries VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *`,
		entry.ID,
		entry.CreatedBy,
		entry.Value,
		entry.Users,
		entry.Split,
		entry.Tags,
		entry.Categories,
		entry.Date); err != nil {
		return fmt.Errorf("error creating thread: %w", err)
	}
	return nil
}

func (s *EntryStore) UpdateEntry(entry *db.Entry) error {
	if err := s.Get(entry, `UPDATE entries SET created_by=$2, value=$3, users=$4, split=$5, tags=$6, categories=$7, date=$8) WHERE id=$1`,
		entry.ID,
		entry.CreatedBy,
		entry.Value,
		entry.Users,
		entry.Split,
		entry.Tags,
		entry.Categories,
		entry.Date); err != nil {
		return fmt.Errorf("error updating entry: %w", err)
	}
	return nil
}

func (s *EntryStore) DeleteEntry(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM entries WHERE id=$1`, id); err != nil {
		return fmt.Errorf("error deleting entry: %w", err)
	}
	return nil
}
