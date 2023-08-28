package repo

import (
	"gorm.io/gorm"

	"github.com/projek-elements/internal/model"
	"github.com/projek-elements/internal/types/trepo"
)

// Session is the session repository.
type Session struct {
	db *gorm.DB
	*ImplementCommonRepository[model.Session]
}

var (
	_ trepo.Session = &Session{}
)

// NewSession creates new Session instance.
func NewSession(db *gorm.DB) *Session {
	return &Session{
		db:                        db,
		ImplementCommonRepository: NewCommon[model.Session](db),
	}
}
