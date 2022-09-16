package ormx

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type DBOption func(db *gorm.DB)

func RowEffect(affect *int64) DBOption {
	return func(db *gorm.DB) {
		*affect = db.RowsAffected
	}
}

type DBXQuery[T any] struct {
	db *gorm.DB
}

func NewDBXQuery[T any](ctx context.Context) *DBXQuery[T] {
	return &DBXQuery[T]{
		db: DB(ctx).Model(new(T)),
	}
}

func (us *DBXQuery[T]) DB() *gorm.DB {
	return us.db
}

func (us *DBXQuery[T]) Table(name string) *DBXQuery[T] {
	us.db = us.db.Table(name)
	return us
}

func (us *DBXQuery[T]) List(opts ...DBOption) ([]*T, error) {
	res := make([]*T, 0)
	err := us.db.Find(&res).Error
	for _, opt := range opts {
		opt(us.db)
	}
	return res, err
}

func (us *DBXQuery[T]) One() (*T, error) {
	res := new(T)
	err := us.db.First(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return res, err
}

func (us *DBXQuery[T]) OneOrFailed() (*T, error) {
	res := new(T)
	err := us.db.First(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return res, err
}

func (us *DBXQuery[T]) DestTo(value any) (any, error) {
	err := us.db.Find(value).Error
	return value, err
}

func (us *DBXQuery[T]) Exist() (bool, error) {
	res := new(T)
	err := us.db.First(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}

func (us *DBXQuery[T]) Count() (int64, error) {
	var count int64
	err := us.db.Model(new(T)).Find(&count).Error
	return count, err
}

func (us *DBXQuery[T]) Limit(limit int) *DBXQuery[T] {
	us.db = us.db.Limit(limit)
	return us
}

func (us *DBXQuery[T]) Offset(offset int) *DBXQuery[T] {
	us.db = us.db.Offset(offset)
	return us
}

func (us *DBXQuery[T]) Order(order string) *DBXQuery[T] {
	us.db = us.db.Order(order)
	return us
}

func (us *DBXQuery[T]) Like(key, value string) *DBXQuery[T] {
	us.db = us.db.Where(fmt.Sprintf("%s like ?", key), "%"+value+"%")
	return us
}

func (us *DBXQuery[T]) EQ(key, value string) *DBXQuery[T] {
	us.db = us.db.Where(key+" = ?", value)
	return us
}

func (us *DBXQuery[T]) NEQ(key, value string) *DBXQuery[T] {
	us.db = us.db.Where(key+" != ?", value)
	return us
}

func (us *DBXQuery[T]) PluckStr(column string) ([]string, error) {
	res := make([]string, 0)
	err := us.db.Pluck(column, &res).Error
	return res, err

}

func (us *DBXQuery[T]) PluckInt(column string) ([]int, error) {
	res := make([]int, 0)
	err := us.db.Pluck(column, &res).Error
	return res, err
}

func (us *DBXQuery[T]) PluckInt64(column string) ([]int64, error) {
	res := make([]int64, 0)
	err := us.db.Pluck(column, &res).Error
	return res, err
}

func (us *DBXQuery[T]) Where(query string, v ...any) *DBXQuery[T] {
	us.db = us.db.Where(query, v...)
	return us
}

func (us *DBXQuery[T]) Select(query string, v ...any) *DBXQuery[T] {
	us.db = us.db.Select(query, v...)
	return us
}

func (us *DBXQuery[T]) Joins(query string, v ...any) *DBXQuery[T] {
	us.db = us.db.Joins(query, v...)
	return us
}

func (us *DBXQuery[T]) IsWhere(b bool, query string, v ...any) *DBXQuery[T] {
	if b {
		us.db = us.db.Where(query, v...)
	}
	return us
}

func (us *DBXQuery[T]) Preload(name string, args ...any) *DBXQuery[T] {
	us.db = us.db.Preload(name, args...)
	return us
}

func (us *DBXQuery[T]) Group(name string) *DBXQuery[T] {
	us.db = us.db.Group(name)
	return us
}
