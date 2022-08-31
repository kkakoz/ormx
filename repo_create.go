package ormx

import (
	"context"
	"gorm.io/gorm"
)

type DBXCreate[T any] struct {
	db *gorm.DB
}

func NewDBXCreate[T any](ctx context.Context) *DBXCreate[T] {
	return &DBXCreate[T]{
		db: DB(ctx),
	}
}

func (us *DBXCreate[T]) DB() *gorm.DB {
	return us.db
}

func (us *DBXCreate[T]) Add(value *T) error {
	err := us.db.Create(value).Error
	return err
}

func (us *DBXCreate[T]) AddList(values []*T) error {
	err := us.db.CreateInBatches(values, 1000).Error
	return err
}
