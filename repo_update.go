package ormx

import (
	"context"
	"gorm.io/gorm"
)

type DBXUpdate[T any] struct {
	db *gorm.DB
}

func NewDBXUpdate[T any](ctx context.Context) *DBXUpdate[T] {
	return &DBXUpdate[T]{
		db: DB(ctx).Model(new(T)),
	}
}

func (us *DBXUpdate[T]) DB() *gorm.DB {
	return us.db
}

func (us *DBXUpdate[T]) Table(name string) *DBXUpdate[T] {
	us.db = us.db.Table(name)
	return us
}

func (us *DBXUpdate[T]) Where(query string, v ...any) *DBXUpdate[T] {
	us.db = us.db.Where(query, v...)
	return us
}

func (us *DBXUpdate[T]) IsWhere(b bool, query string, v ...any) *DBXUpdate[T] {
	if b {
		us.db = us.db.Where(query, v...)
	}
	return us
}

func (us *DBXUpdate[T]) Update(column string, value any) error {
	return us.db.Update(column, value).Error
}

func (us *DBXUpdate[T]) Updates(value T) error {
	return us.db.Updates(value).Error
}

func (us *DBXUpdate[T]) UpdatesMap(value map[string]any) error {
	return us.db.Updates(value).Error
}
