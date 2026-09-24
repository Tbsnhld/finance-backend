package postgres

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
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
	ID         uuid.UUID  `db:"id"`
	CreatedBy  User       `db:"createdBy"`
	Value      float32    `db:"value"`
	Users      []User     `db:"users"`
	Split      []float32  `db:"split"`
	Tags       []Tag      `db:"tags"`
	Categories []Category `db:"categories"`
	Date       time.Time  `db:"date"`
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
	Tags() ([]Tags, error)
	CreateTag(tag *Tag) error
	UpdateTag(tag *Tag) error
	DeleteTag(id uuid.UUID) error
}
