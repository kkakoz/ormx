package ormx

import (
	"context"
	"gorm.io/gorm"
)

type DBXDelete[T any] struct {
	db *gorm.DB
}

func NewDBXDelete[T any](ctx context.Context) *DBXDelete[T] {
	return &DBXDelete[T]{
		db: DB(ctx).Model(new(T)),
	}
}

func (us *DBXDelete[T]) DB() *gorm.DB {
	return us.db
}

func (us *DBXDelete[T]) Table(name string) *DBXDelete[T] {
	us.db = us.db.Table(name)
	return us
}

func (us *DBXDelete[T]) Where(query string, v ...any) *DBXDelete[T] {
	us.db = us.db.Where(query, v...)
	return us
}

func (us *DBXDelete[T]) IsWhere(b bool, query string, v ...any) *DBXDelete[T] {
	if b {
		us.db = us.db.Where(query, v...)
	}
	return us
}

func (us *DBXDelete[T]) Exec(opts ...DBOption) error {
	t := new(T)
	err := us.db.Delete(t).Error
	for _, opt := range opts {
		opt(us.db)
	}
	return err
}
