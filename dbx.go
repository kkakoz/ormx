package ormx

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type DBX[T any] struct {
	db *gorm.DB
}

func NewDBX[T any](ctx context.Context) *DBX[T] {
	return &DBX[T]{
		db: DB(ctx),
	}
}

func (us *DBX[T]) Table(name string) *DBX[T] {
	us.db.Table(name)
	return us
}

func (us *DBX[T]) List() ([]*T, error) {
	res := make([]*T, 0)
	err := us.db.Find(&res).Error
	return res, err
}

func (us *DBX[T]) One() (*T, error) {
	res := new(T)
	err := us.db.First(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return res, err
}

func (us *DBX[T]) OneOrFailed() (*T, error) {
	res := new(T)
	err := us.db.First(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return res, err
}

func (us *DBX[T]) DestTo(value any) (any, error) {
	err := us.db.Find(value).Error
	return value, err
}

func (us *DBX[T]) Exist() (bool, error) {
	res := new(T)
	err := us.db.First(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}

func (us *DBX[T]) Count() (int64, error) {
	var count int64
	err := us.db.Model(new(T)).Find(&count).Error
	return count, err
}

func (us *DBX[T]) Limit(limit int) *DBX[T] {
	us.db.Limit(limit)
	return us
}

func (us *DBX[T]) Offset(offset int) *DBX[T] {
	us.db.Offset(offset)
	return us
}

func (us *DBX[T]) Pluck(column string, dest any) error {
	return us.db.Pluck(column, dest).Error
}

func (us *DBX[T]) Where(query string, v ...any) *DBX[T] {
	us.db = us.db.Where(query, v...)
	return us
}

func (us *DBX[T]) Select(query string, v ...any) *DBX[T] {
	us.db = us.db.Select(query, v...)
	return us
}

func (us *DBX[T]) Joins(query string, v ...any) *DBX[T] {
	us.db = us.db.Joins(query, v...)
	return us
}

func (us *DBX[T]) IsWhere(b bool, query string, v ...any) *DBX[T] {
	if b {
		us.db = us.db.Where(query, v...)
	}
	return us
}

func (us *DBX[T]) Update(column string, value any) error {
	return us.db.Update(column, value).Error
}

func (us *DBX[T]) Updates(value T) error {
	return us.db.Updates(value).Error
}

func (us *DBX[T]) UpdatesMap(value map[string]any) error {
	return us.db.Updates(value).Error
}

func (us *DBX[T]) Delete() error {
	t := new(T)
	err := us.db.Delete(t).Error
	return err
}

func (us *DBX[T]) Preload(name string, args ...any) {
	us.db = us.db.Preload(name, args...)
}
