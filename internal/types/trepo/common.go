package trepo

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/projek-elements/pkg/util"
)

// Common represents the common repository.
type Common[Model any] interface {
	// DoTx create a transaction.
	DoTx(func(tx *gorm.DB) error) error
	// Create a model.
	Create(c context.Context, m *Model) (*Model, error)
	// CreateMany creates models by a list.
	CreateMany(c context.Context, ms ...*Model) ([]*Model, error)
	// NOTE: BatchInsert required gorm version >= 2
	BatchInsert(c context.Context, ms ...*Model) ([]*Model, error)
	// Find model by ID.
	FindByID(c context.Context, id int) (*Model, error)
	// Delete model by ID.
	DeleteByID(c context.Context, id int) error
	// Delete must be provided condition. Delete(c context.Context, cond any, args ...any) error
	UpdateByID(c context.Context, updates UpdationData[Model], id int) error
	// Update must be provided condition.
	Update(c context.Context, updates UpdationData[Model], cond any, args ...any) error
}

// CreateImplFn creates new repository handler function
type CreateImplFn[Impl any] func(*gorm.DB) *Impl

// UpdationData represents the update model of repository.
type UpdationData[M any] interface {
	// Model creates new instance of M.
	Model() *M
}

// Model is the common model.
type Model[T any] struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	updationData[T]
}

type updationData[M any] struct{}

// Model implements UpdationData.
func (*updationData[M]) Model() *M {
	return new(M)
}

var (
	_ UpdationData[util.Something] = &updationData[util.Something]{}
)

// NewOrTxFn Creates or extends transaction.
type NewOrTxFn[M any] func() *M

type NewWithDB[M any] func(db *gorm.DB) *M
