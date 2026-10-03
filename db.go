package main

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	Title     string    `db:"title"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	UserName  string    `db:"user_name"`
}

type Tag struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

type Category struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

type Entry struct {
	ID         uuid.UUID   `db:"id"`
	CreatedBy  uuid.UUID   `db:"created_by"`
	Value      float32     `db:"value"`
	Users      []uuid.UUID `db:"users"`
	Split      []float32   `db:"split"`
	Tags       []uuid.UUID `db:"tags"`
	Categories uuid.UUID   `db:"categories"`
	Date       time.Time   `db:"date"`
}

type EntryStore interface {
	Entry(id uuid.UUID) (Entry, error)
	Entries() ([]Entry, error)
	EntriesByUser(userId uuid.UUID) ([]Entry, error)
	EntriesByCategory(categoryId uuid.UUID) ([]Entry, error)
	EntriesByTag(tagId uuid.UUID) ([]Entry, error)
	CreateEntry(entry *Entry) error
	UpdateEntry(entry *Entry) error
	DeleteEntry(id uuid.UUID) error
}

type UserStore interface {
	User(id uuid.UUID) (User, error)
	Users() ([]User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id uuid.UUID) error
}

type TagStore interface {
	Tag(id uuid.UUID) (Tag, error)
	Tags() ([]Tag, error)
	CreateTag(tag *Tag) error
	UpdateTag(tag *Tag) error
	DeleteTag(id uuid.UUID) error
}

type CategoryStore interface {
	Category(id uuid.UUID) (Category, error)
	Categories() ([]Category, error)
	CreateCategory(tag *Category) error
	UpdateCategory(tag *Category) error
	DeleteCategory(id uuid.UUID) error
}

type Store interface {
	EntryStore
	UserStore
	TagStore
	CategoryStore
}
