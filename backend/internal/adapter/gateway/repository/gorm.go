package repository

import (
	"context"

	"github.com/schema-creator/schema-creator/schema-creator/internal/usecase/dai"
	"gorm.io/gorm"
)

type GormRepo struct {
	gorm *gorm.DB
}

func NewGormRepo(gorm *gorm.DB) *GormRepo {
	return &GormRepo{
		gorm: gorm,
	}
}

var _ dai.DataAccessInterfaces = (*GormRepo)(nil)

func (r *GormRepo) Transaction(ctx context.Context, fn func(context.Context, dai.DataAccessInterfaces) error) error {
	err := r.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := fn(ctx, NewGormRepo(tx)); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
