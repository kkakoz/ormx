package ormx

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type QueryOption func(db *gorm.DB)

func RowEffect(affect *int64) QueryOption {
	return func(db *gorm.DB) {
		*affect = db.RowsAffected
	}
}

type IDBExec[T any] interface {
	Where(query string, v ...any) *T
}

type DBXQuery[T any, Query any] struct {
	db    *gorm.DB
	query *Query
}

func NewDBXQuery[T any, Query any](ctx context.Context, query *Query) *DBXQuery[T, Query] {
	return &DBXQuery[T, Query]{
		db:    DB(ctx),
		query: query,
	}
}

func (us *DBXQuery[T, Query]) DB() *gorm.DB {
	return us.db
}

func (us *DBXQuery[T, Query]) Table(name string) *DBXQuery[T, Query] {
	us.db.Table(name)
	return us
}

func (us *DBXQuery[T, Query]) List(opts ...QueryOption) ([]*T, error) {
	res := make([]*T, 0)
	err := us.db.Find(&res).Error
	for _, opt := range opts {
		opt(us.db)
	}
	return res, err
}

func (us *DBXQuery[T, Query]) One() (*T, error) {
	res := new(T)
	err := us.db.First(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return res, err
}

func (us *DBXQuery[T, Query]) OneOrFailed() (*T, error) {
	res := new(T)
	err := us.db.First(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return res, err
}

func (us *DBXQuery[T, Query]) DestTo(value any) (any, error) {
	err := us.db.Find(value).Error
	return value, err
}

func (us *DBXQuery[T, Query]) Exist() (bool, error) {
	res := new(T)
	err := us.db.First(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}

func (us *DBXQuery[T, Query]) Count() (int64, error) {
	var count int64
	err := us.db.Model(new(T)).Find(&count).Error
	return count, err
}

func (us *DBXQuery[T, Query]) Limit(limit int) *Query {
	us.db.Limit(limit)
	return us.query
}

func (us *DBXQuery[T, Query]) Offset(offset int) *Query {
	us.db.Offset(offset)
	return us.query
}

func (us *DBXQuery[T, Query]) Pluck(column string, dest any) error {
	return us.db.Pluck(column, dest).Error
}

func (us *DBXQuery[T, Query]) Where(query string, v ...any) *Query {
	us.db = us.db.Where(query, v...)
	return us.query
}

func (us *DBXQuery[T, Query]) Select(query string, v ...any) *Query {
	us.db = us.db.Select(query, v...)
	return us.query
}

func (us *DBXQuery[T, Query]) Joins(query string, v ...any) *Query {
	us.db = us.db.Joins(query, v...)
	return us.query
}

func (us *DBXQuery[T, Query]) IsWhere(b bool, query string, v ...any) *Query {
	if b {
		us.db = us.db.Where(query, v...)
	}
	return us.query
}

func (us *DBXQuery[T, Query]) Preload(name string, args ...any) *Query {
	us.db = us.db.Preload(name, args...)
	return us.query
}

func (us *DBXQuery[T, Query]) Group(name string) *Query {
	us.db = us.db.Group(name)
	return us.query
}
