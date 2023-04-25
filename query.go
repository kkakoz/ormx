package ormx

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type defaultQuery[T any] struct {
	db *gorm.DB
}

func NewQuery[T any](db *gorm.DB) defaultQuery[T] {
	return defaultQuery[T]{db: db}
}

func (us *defaultQuery[T]) Table(name string) *defaultQuery[T] {
	us.db = us.db.Table(name)
	return us
}

func (us *defaultQuery[T]) List(opts ...DBOption) ([]*T, error) {
	res := make([]*T, 0)
	err := us.db.Find(&res).Error
	for _, opt := range opts {
		opt(us.db)
	}
	return res, err
}

func (us *defaultQuery[T]) One() (*T, error) {
	res := new(T)
	err := us.db.First(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return res, err
}

func (us *defaultQuery[T]) OneOrFailed() (*T, error) {
	res := new(T)
	err := us.db.First(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return res, err
}

func (us *defaultQuery[T]) DestTo(value any) (any, error) {
	err := us.db.Find(value).Error
	return value, err
}

func (us *defaultQuery[T]) Exist() (bool, error) {
	res := new(T)
	err := us.db.First(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}

func (us *defaultQuery[T]) Count() (int64, error) {
	var count int64
	err := us.db.Model(new(T)).Find(&count).Error
	return count, err
}

func (us *defaultQuery[T]) Limit(limit int) *defaultQuery[T] {
	us.db = us.db.Limit(limit)
	return us
}

func (us *defaultQuery[T]) Offset(offset int) *defaultQuery[T] {
	us.db = us.db.Offset(offset)
	return us
}

func (us *defaultQuery[T]) Order(order string) *defaultQuery[T] {
	us.db = us.db.Order(order)
	return us
}

func (us *defaultQuery[T]) Like(key, value string) *defaultQuery[T] {
	us.db = us.db.Where(fmt.Sprintf("%s like ?", key), "%"+value+"%")
	return us
}

func (us *defaultQuery[T]) EQ(key, value string) *defaultQuery[T] {
	us.db = us.db.Where(key+" = ?", value)
	return us
}

func (us *defaultQuery[T]) NEQ(key, value string) *defaultQuery[T] {
	us.db = us.db.Where(key+" != ?", value)
	return us
}

func (us *defaultQuery[T]) PluckStr(column string) ([]string, error) {
	res := make([]string, 0)
	err := us.db.Pluck(column, &res).Error
	return res, err
}

func (us *defaultQuery[T]) PluckInt(column string) ([]int, error) {
	res := make([]int, 0)
	err := us.db.Pluck(column, &res).Error
	return res, err
}

func (us *defaultQuery[T]) PluckInt64(column string) ([]int64, error) {
	res := make([]int64, 0)
	err := us.db.Pluck(column, &res).Error
	return res, err
}

func (us *defaultQuery[T]) Where(query string, v ...any) *defaultQuery[T] {
	us.db = us.db.Where(query, v...)
	return us
}

func (us *defaultQuery[T]) Select(query string, v ...any) *defaultQuery[T] {
	us.db = us.db.Select(query, v...)
	return us
}

func (us *defaultQuery[T]) Joins(query string, v ...any) *defaultQuery[T] {
	us.db = us.db.Joins(query, v...)
	return us
}

func (us *defaultQuery[T]) IsWhere(b bool, query string, v ...any) *defaultQuery[T] {
	if b {
		us.db = us.db.Where(query, v...)
	}
	return us
}

func (us *defaultQuery[T]) Preload(name string, args ...any) *defaultQuery[T] {
	us.db = us.db.Preload(name, args...)
	return us
}

func (us *defaultQuery[T]) Group(name string) *defaultQuery[T] {
	us.db = us.db.Group(name)
	return us
}
