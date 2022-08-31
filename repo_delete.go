package ormx

import (
	"context"
	"gorm.io/gorm"
)

type DBXDelete[T any, Query any] struct {
	db    *gorm.DB
	query *Query
}

func NewDBXDelete[T any, Query any](ctx context.Context, query *Query) *DBXDelete[T, Query] {
	return &DBXDelete[T, Query]{
		db:    DB(ctx),
		query: query,
	}
}

func (us *DBXDelete[T, Query]) DB() *gorm.DB {
	return us.db
}

func (us *DBXDelete[T, Query]) Table(name string) *DBXDelete[T, Query] {
	us.db.Table(name)
	return us
}

func (us *DBXDelete[T, Query]) Where(query string, v ...any) *Query {
	us.db = us.db.Where(query, v...)
	return us.query
}

func (us *DBXDelete[T, Query]) IsWhere(b bool, query string, v ...any) *Query {
	if b {
		us.db = us.db.Where(query, v...)
	}
	return us.query
}

func (us *DBXDelete[T, Query]) Delete() error {
	t := new(T)
	err := us.db.Delete(t).Error
	return err
}
