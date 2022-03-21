package ormx

import (
	"context"

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
	Updates(ctx context.Context, value T, opts ...opts.Option) error
}

type Repo[T any] struct {
	errHandle ErrHandler
}

type Option[T any] func(r *Repo[T])

func NewRepo[T any](opts ...Option[T]) *Repo[T] {
	r := &Repo[T]{
		errHandle: DefaultErrHandler,
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func (r *Repo[T]) Add(ctx context.Context, value *T) error {
	db := DB(ctx)
	err := db.Create(value).Error
	return err
}

func (r *Repo[T]) AddList(ctx context.Context, value []*T) error {
	db := DB(ctx)
	err := db.CreateInBatches(value, 1000).Error
	return r.errHandle(err)
}

func (r *Repo[T]) Get(ctx context.Context, opts ...opts.Option) (*T, error) {
	db := DB(ctx)
	target := new(T)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.First(target).Error
	return target, r.errHandle(err)
}

func (r *Repo[T]) GetById(ctx context.Context, id any) (*T, error) {
	db := DB(ctx)
	target := new(T)
	err := db.First(target, id).Error
	return target, r.errHandle(err)
}

func (r *Repo[T]) GetExist(ctx context.Context, opts ...opts.Option) (bool, error) {
	db := DB(ctx)
	target := new(T)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.First(target).Error
	if err != nil {
		return false, r.errHandle(err)
	}
	return true, r.errHandle(err)
}

func (r *Repo[T]) GetList(ctx context.Context, opts ...opts.Option) ([]*T, error) {
	db := DB(ctx)
	list := make([]*T, 0)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Find(&list).Error
	return list, r.errHandle(err)
}

func (r *Repo[T]) Count(ctx context.Context, opts ...opts.Option) (int64, error) {
	db := DB(ctx)
	var count int64
	target := new(T)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Model(target).Count(&count).Error
	return count, r.errHandle(err)
}

func (r *Repo[T]) DeleteById(ctx context.Context, id any) error {
	db := DB(ctx)
	err := db.Delete(new(T), id).Error
	return r.errHandle(err)
}

func (r *Repo[T]) Delete(ctx context.Context, opts ...opts.Option) error {
	db := DB(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Delete(new(T)).Error
	return r.errHandle(err)
}

func (r *Repo[T]) Updates(ctx context.Context, value any, opts ...opts.Option) error {
	db := DB(ctx)
	target := new(T)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Model(target).Updates(value).Error
	return r.errHandle(err)
}
