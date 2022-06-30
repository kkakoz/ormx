package ormx

import (
	"context"
	"errors"
	"gorm.io/gorm"

	"github.com/kkakoz/ormx/opt"
)

type IRepo[T any] interface {
	Add(ctx context.Context, value *T) error
	AddList(ctx context.Context, value []*T) error
	Get(ctx context.Context, opts ...opt.Option) (*T, error)
	Pluck(ctx context.Context, column string, slice any, opts ...opt.Option) error
	GetById(ctx context.Context, conds ...any) (*T, error)
	GetExist(ctx context.Context, opts ...opt.Option) (bool, error)
	GetList(ctx context.Context, opts ...opt.Option) ([]*T, error)
	PageList(ctx context.Context, offset, limit int, opts ...opt.Option) ([]*T, int64, error)
	Count(ctx context.Context, opts ...opt.Option) (int64, error)
	DeleteById(ctx context.Context, id any) error
	Delete(ctx context.Context, opts ...opt.Option) error
	Updates(ctx context.Context, value any, opts ...opt.Option) error
}

type repo[T any] struct {
	errHandle ErrHandler
}

func NewRepo[T any](opts ...Option[T]) IRepo[T] {
	r := &repo[T]{
		errHandle: DefaultErrHandler,
	}
	for _, o := range opts {
		o(r)
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

func (r *repo[T]) Get(ctx context.Context, opts ...opt.Option) (*T, error) {
	db := DB(ctx)
	target := new(T)
	db = opt.OptionsDB(db, opts...)
	err := db.First(target).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return target, r.errHandle(err)
}

func (r *repo[T]) Pluck(ctx context.Context, column string, slice any, opts ...opt.Option) error {
	db := DB(ctx)
	db = opt.OptionsDB(db, opts...)
	return r.errHandle(db.Pluck(column, slice).Error)
}

func (r *repo[T]) GetById(ctx context.Context, conds ...any) (*T, error) {
	db := DB(ctx)
	target := new(T)
	err := db.First(target, conds...).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return target, r.errHandle(err)
}

func (r *repo[T]) GetExist(ctx context.Context, opts ...opt.Option) (bool, error) {
	db := DB(ctx)
	target := new(T)
	db = opt.OptionsDB(db, opts...)
	err := db.First(target).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, r.errHandle(err)
	}
	return true, r.errHandle(err)
}

func (r *repo[T]) GetList(ctx context.Context, opts ...opt.Option) ([]*T, error) {
	db := DB(ctx)
	list := make([]*T, 0)
	db = opt.OptionsDB(db, opts...)
	err := db.Find(&list).Error
	return list, r.errHandle(err)
}

func (r *repo[T]) PageList(ctx context.Context, offset, limit int, opts ...opt.Option) ([]*T, int64, error) {
	db := DB(ctx)
	list := make([]*T, 0)
	db = opt.OptionsDB(db, opts...)
	t := new(T)
	var count int64
	db.Model(t).Count(&count)
	err := db.Limit(limit).Offset(offset).Find(&list).Error
	return list, count, r.errHandle(err)
}

func (r *repo[T]) Count(ctx context.Context, opts ...opt.Option) (int64, error) {
	db := DB(ctx)
	var count int64
	target := new(T)
	db = opt.OptionsDB(db, opts...)
	err := db.Model(target).Count(&count).Error
	return count, r.errHandle(err)
}

func (r *repo[T]) DeleteById(ctx context.Context, id any) error {
	db := DB(ctx)
	err := db.Delete(new(T), id).Error
	return r.errHandle(err)
}

func (r *repo[T]) Delete(ctx context.Context, opts ...opt.Option) error {
	db := DB(ctx)
	db = opt.OptionsDB(db, opts...)
	err := db.Delete(new(T)).Error
	return r.errHandle(err)
}

func (r *repo[T]) Updates(ctx context.Context, value any, opts ...opt.Option) error {
	db := DB(ctx)
	target := new(T)
	db = opt.OptionsDB(db, opts...)
	err := db.Model(target).Updates(value).Error
	return r.errHandle(err)
}
