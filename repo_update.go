package ormx

import (
	"context"
	"gorm.io/gorm"
)

type DBXUpdate[T any, Query any] struct {
	db    *gorm.DB
	query *Query
}

func NewDBXUpdate[T any, Query any](ctx context.Context, query *Query) *DBXUpdate[T, Query] {
	return &DBXUpdate[T, Query]{
		db:    DB(ctx),
		query: query,
	}
}

func (us *DBXUpdate[T, Query]) DB() *gorm.DB {
	return us.db
}

func (us *DBXUpdate[T, Query]) Table(name string) *DBXUpdate[T, Query] {
	us.db.Table(name)
	return us
}

func (us *DBXUpdate[T, Query]) Where(query string, v ...any) *Query {
	us.db = us.db.Where(query, v...)
	return us.query
}

func (us *DBXUpdate[T, Query]) IsWhere(b bool, query string, v ...any) *Query {
	if b {
		us.db = us.db.Where(query, v...)
	}
	return us.query
}

func (us *DBXUpdate[T, Query]) Update(column string, value any) error {
	return us.db.Model(new(T)).Update(column, value).Error
}

func (us *DBXUpdate[T, Query]) Updates(value T) error {
	return us.db.Model(new(T)).Updates(value).Error
}

func (us *DBXUpdate[T, Query]) UpdatesMap(value map[string]any) error {
	return us.db.Model(new(T)).Updates(value).Error
}
