package postgres

import (
	"finance-backend/db"
	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	*sqlx.DB
}

func NewUserStore(db *sqlx.DB) *UserStore {
	return &UserStore{
		DB: db,
	}
}

func (s *UserStore) User(id uuid.UUID) (db.User, error) {
	var user db.User
	if err := s.Get(&user, `SELECT * FROM users WHERE id=$1`, id); err != nil {
		return db.User{}, fmt.Errorf("error getting user: %w", err)
	}
	return user, nil
}

func (s *UserStore) Users() ([]db.User, error) {
	var users []db.User
	if err := s.Select(&users, `SELECT * FROM users`); err != nil {
		return []db.User{}, fmt.Errorf("error getting users: %w", err)
	}
	return users, nil
}

func (s *UserStore) CreateUser(user *db.User) error {
	if err := s.Get(user, `INSERT INTO users VALUES ($1, $2, $3, $4, $5) RETURNING *`,
		user.ID,
		user.FirstName,
		user.LastName,
		user.UserName,
		user.Title,
	); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func (s *UserStore) UpdateUser(user *db.User) error {
	if err := s.Get(user, `UPDATE users SET first_name=$2, last_name=$3, user_name=$4, title=$5 ) WHERE id=$1`,
		user.ID,
		user.FirstName,
		user.LastName,
		user.UserName,
		user.Title,
	); err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

func (s *UserStore) DeleteUser(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM users WHERE id=$1`, id); err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}
