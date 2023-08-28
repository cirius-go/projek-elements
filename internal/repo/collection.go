package repo

import (
	"gorm.io/gorm"
)

// Collection contains multiple repositories.
type Collection struct {
	db      *gorm.DB
	Session *Session
}

func NewCollection(db *gorm.DB) *Collection {
	c := &Collection{
		Session: NewSession(db),
	}

	return c
}
