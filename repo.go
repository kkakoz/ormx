package ormx

import (
	"context"
	"errors"
	"gorm.io/gorm"

	"github.com/kkakoz/ormx/opts"
)

type IRepo[T any] interface {
	Add(ctx context.Context, value *T) error
	AddList(ctx context.Context, value []*T) error
	Get(ctx context.Context, opts ...opts.Option) (*T, error)
	GetById(ctx context.Context, id any) (*T, error)
	GetExist(ctx context.Context, opts ...opts.Option) (bool, error)
	GetList(ctx context.Context, opts ...opts.Option) ([]*T, error)
	Count(ctx context.Context, opts ...opts.Option) (int64, error)
	DeleteById(ctx context.Context, id any) error
	Delete(ctx context.Context, opts ...opts.Option) error
	Updates(ctx context.Context, value any, opts ...opts.Option) error
}

type repo[T any] struct {
	errHandle ErrHandler
}

type Option[T any] func(r *repo[T])

func NewRepo[T any](opts ...Option[T]) IRepo[T] {
	r := &repo[T]{
		errHandle: DefaultErrHandler,
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func (r *repo[T]) Add(ctx context.Context, value *T) error {
	db := DB(ctx)
	err := db.Create(value).Error
	return err
}

func (r *repo[T]) AddList(ctx context.Context, value []*T) error {
	db := DB(ctx)
	err := db.CreateInBatches(value, 1000).Error
	return r.errHandle(err)
}

func (r *repo[T]) Get(ctx context.Context, opts ...opts.Option) (*T, error) {
	db := DB(ctx)
	target := new(T)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.First(target).Error
	return target, r.errHandle(err)
}

func (r *repo[T]) GetById(ctx context.Context, id any) (*T, error) {
	db := DB(ctx)
	target := new(T)
	err := db.First(target, id).Error
	return target, r.errHandle(err)
}

func (r *repo[T]) GetExist(ctx context.Context, opts ...opts.Option) (bool, error) {
	db := DB(ctx)
	target := new(T)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.First(target).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, r.errHandle(err)
	}
	return true, r.errHandle(err)
}

func (r *repo[T]) GetList(ctx context.Context, opts ...opts.Option) ([]*T, error) {
	db := DB(ctx)
	list := make([]*T, 0)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Find(&list).Error
	return list, r.errHandle(err)
}

func (r *repo[T]) Count(ctx context.Context, opts ...opts.Option) (int64, error) {
	db := DB(ctx)
	var count int64
	target := new(T)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Model(target).Count(&count).Error
	return count, r.errHandle(err)
}

func (r *repo[T]) DeleteById(ctx context.Context, id any) error {
	db := DB(ctx)
	err := db.Delete(new(T), id).Error
	return r.errHandle(err)
}

func (r *repo[T]) Delete(ctx context.Context, opts ...opts.Option) error {
	db := DB(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Delete(new(T)).Error
	return r.errHandle(err)
}

func (r *repo[T]) Updates(ctx context.Context, value any, opts ...opts.Option) error {
	db := DB(ctx)
	target := new(T)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Model(target).Updates(value).Error
	return r.errHandle(err)
}
