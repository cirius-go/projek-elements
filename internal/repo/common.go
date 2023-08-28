package repo

import (
	"context"

	"gorm.io/gorm"

	"github.com/projek-elements/internal/types/trepo"
	"github.com/projek-elements/pkg/util"
)

type ImplementCommonRepository[Model any] struct {
	db *gorm.DB
}

// BatchInsert implements trepo.CommonRepository.
func (*ImplementCommonRepository[Model]) BatchInsert(c context.Context, ms ...*Model) ([]*Model, error) {
	panic("unimplemented")
}

// Create implements trepo.CommonRepository.
func (*ImplementCommonRepository[Model]) Create(c context.Context, m *Model) (*Model, error) {
	panic("unimplemented")
}

// CreateMany implements trepo.CommonRepository.
func (*ImplementCommonRepository[Model]) CreateMany(c context.Context, ms ...*Model) ([]*Model, error) {
	panic("unimplemented")
}

// DeleteByID implements trepo.CommonRepository.
func (*ImplementCommonRepository[Model]) DeleteByID(c context.Context, id int) error {
	panic("unimplemented")
}

// DoTx implements trepo.CommonRepository.
func (*ImplementCommonRepository[Model]) DoTx(func(tx *gorm.DB) error) error {
	panic("unimplemented")
}

// FindByID implements trepo.CommonRepository.
func (*ImplementCommonRepository[Model]) FindByID(c context.Context, id int) (*Model, error) {
	panic("unimplemented")
}

// Update implements trepo.CommonRepository.
func (*ImplementCommonRepository[Model]) Update(c context.Context, updates trepo.UpdationData[Model], cond any, args ...any) error {
	panic("unimplemented")
}

// UpdateByID implements trepo.CommonRepository.
func (*ImplementCommonRepository[Model]) UpdateByID(c context.Context, updates trepo.UpdationData[Model], id int) error {
	panic("unimplemented")
}

var (
	_ trepo.Common[util.Something] = &ImplementCommonRepository[util.Something]{}
)

// NewCommon creates new common repository.
func NewCommon[M any](db *gorm.DB) *ImplementCommonRepository[M] {
	return &ImplementCommonRepository[M]{
		db: db,
	}
}
