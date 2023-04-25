package ormx

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type WithQuery[ResQuery any, T any] struct {
	db  *gorm.DB
	res *ResQuery
}

func NewWithQuery[Query any, T any](query *Query, db *gorm.DB) *WithQuery[Query, T] {
	return &WithQuery[Query, T]{res: query, db: db}
}

func (us *WithQuery[ResQuery, T]) Table(name string) *ResQuery {
	us.db = us.db.Table(name)
	return us.res
}

func (us *WithQuery[ResQuery, T]) List() ([]*T, error) {
	res := make([]*T, 0)
	err := us.db.Find(&res).Error
	return res, err
}

func (us *WithQuery[ResQuery, T]) One() (*T, error) {
	res := new(T)
	err := us.db.First(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return res, err
}

func (us *WithQuery[ResQuery, T]) OneOrFailed() (*T, error) {
	res := new(T)
	err := us.db.First(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return res, err
}

func (us *WithQuery[ResQuery, T]) DestTo(value any) (any, error) {
	err := us.db.Find(value).Error
	return value, err
}

func (us *WithQuery[ResQuery, T]) Exist() (bool, error) {
	res := new(T)
	err := us.db.First(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}

func (us *WithQuery[ResQuery, T]) Count() (int64, error) {
	var count int64
	err := us.db.Model(new(T)).Find(&count).Error
	return count, err
}

func (us *WithQuery[ResQuery, T]) Limit(limit int) *ResQuery {
	us.db = us.db.Limit(limit)
	return us.res
}

func (us *WithQuery[ResQuery, T]) Offset(offset int) *ResQuery {
	us.db = us.db.Offset(offset)
	return us.res
}

func (us *WithQuery[ResQuery, T]) Order(order string) *ResQuery {
	us.db = us.db.Order(order)
	return us.res
}

func (us *WithQuery[ResQuery, T]) Like(key, value string) *ResQuery {
	us.db = us.db.Where(fmt.Sprintf("%s like ?", key), "%"+value+"%")
	return us.res
}

func (us *WithQuery[ResQuery, T]) EQ(key, value string) *ResQuery {
	us.db = us.db.Where(key+" = ?", value)
	return us.res
}

func (us *WithQuery[ResQuery, T]) NEQ(key, value string) *ResQuery {
	us.db = us.db.Where(key+" != ?", value)
	return us.res
}

func (us *WithQuery[ResQuery, T]) PluckStr(column string) ([]string, error) {
	res := make([]string, 0)
	err := us.db.Pluck(column, &res).Error
	return res, err
}

func (us *WithQuery[ResQuery, T]) PluckInt(column string) ([]int, error) {
	res := make([]int, 0)
	err := us.db.Pluck(column, &res).Error
	return res, err
}

func (us *WithQuery[ResQuery, T]) PluckInt64(column string) ([]int64, error) {
	res := make([]int64, 0)
	err := us.db.Pluck(column, &res).Error
	return res, err
}

func (us *WithQuery[ResQuery, T]) Where(query string, v ...any) *ResQuery {
	us.db = us.db.Where(query, v...)
	return us.res
}

func (us *WithQuery[ResQuery, T]) Select(query string, v ...any) *ResQuery {
	us.db = us.db.Select(query, v...)
	return us.res
}

func (us *WithQuery[ResQuery, T]) Joins(query string, v ...any) *ResQuery {
	us.db = us.db.Joins(query, v...)
	return us.res
}

func (us *WithQuery[ResQuery, T]) IsWhere(b bool, query string, v ...any) *ResQuery {
	if b {
		us.db = us.db.Where(query, v...)
	}
	return us.res
}

func (us *WithQuery[ResQuery, T]) Preload(name string, args ...any) *ResQuery {
	us.db = us.db.Preload(name, args...)
	return us.res
}

func (us *WithQuery[ResQuery, T]) Group(name string) *ResQuery {
	us.db = us.db.Group(name)
	return us.res
}

func (us *WithQuery[ResQuery, T]) Update(column string, value any) error {
	return us.db.Update(column, value).Error
}

func (us *WithQuery[ResQuery, T]) Delete(column string, value any) error {
	return us.db.Update(column, value).Error
}

type UserQuery struct {
	*WithQuery[UserQuery, User]
}

type User struct {
}

func (u UserQuery) IDEQ(id int) *UserQuery {
	return u.WithQuery.Where("id = ?", id)
}

func (u UserQuery) NameEQ(name string) *UserQuery {
	return u.WithQuery.Where("name = ?", name)
}

func UserQuerys() {
	query := &UserQuery{}
	query.WithQuery = NewWithQuery[UserQuery, User](query, &gorm.DB{})

}
