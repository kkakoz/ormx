package repo

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/example/model"
)

type _userQuery[Query any] struct {
	ormx.IDBExec[Query]
	parent *Query
}

type userQuery struct {
	*_userQuery[userQuery]
	*ormx.DBXQuery[model.User, userQuery]
}

type userUpdate struct {
	*_userQuery[userUpdate]
	*ormx.DBXUpdate[model.User, userUpdate]
}

type userDelete struct {
	*_userQuery[userDelete]
	*ormx.DBXDelete[model.User, userDelete]
}

func NewuserQuery(ctx context.Context) *userQuery {
	query := &userQuery{}
	query._userQuery = &_userQuery[userQuery]{}
	query.DBXQuery = ormx.NewDBXQuery[model.User, userQuery](ctx, query)
	query.IDBExec = query.DBXQuery
	query.parent = query
	return query
}

func NewuserUpdate(ctx context.Context) *userUpdate {
	update := &userUpdate{}
	update._userQuery = &_userQuery[userUpdate]{}
	update.DBXUpdate = ormx.NewDBXUpdate[model.User, userUpdate](ctx, update)
	update.IDBExec = update.DBXUpdate
	update.parent = update
	return update
}

func NewuserDelete(ctx context.Context) *userDelete {
	del := &userDelete{}
	del._userQuery = &_userQuery[userDelete]{}
	del.DBXDelete = ormx.NewDBXDelete[model.User, userDelete](ctx, del)
	del.IDBExec = del.DBXDelete
	del.parent = del
	return del
}

func NewuserCreate(ctx context.Context) *ormx.DBXCreate[model.User] {
	return ormx.NewDBXCreate[model.User](ctx)
}

func (us *_userQuery[Query]) ID(id uint) *Query {
	us.IDBExec.Where("id = ?", id)
	return us.parent
}

func (us *_userQuery[Query]) IDGT(id uint) *Query {
	us.IDBExec.Where("id > ?", id)
	return us.parent
}

func (us *_userQuery[Query]) IDLT(id uint) *Query {
	us.IDBExec.Where("id < ?", id)
	return us.parent
}

func (us *_userQuery[Query]) IDLTE(id uint) *Query {
	us.IDBExec.Where("id <= ?", id)
	return us.parent
}

func (us *_userQuery[Query]) IDGTE(id uint) *Query {
	us.IDBExec.Where("id >= ?", id)
	return us.parent
}

func (us *_userQuery[Query]) IDIn(ids []uint) *Query {
	us.IDBExec.Where("id in ?", ids)
	return us.parent
}

func (us *_userQuery[Query]) IDNEq(id uint) *Query {
	us.IDBExec.Where("id != ?", id)
	return us.parent
}

func (us *_userQuery[Query]) IDNotIn(ids []uint) *Query {
	us.IDBExec.Where("id not in ?", ids)
	return us.parent
}

func (us *_userQuery[Query]) Name(name string) *Query {
	us.IDBExec.Where("name = ?", name)
	return us.parent
}

func (us *_userQuery[Query]) NameGT(name string) *Query {
	us.IDBExec.Where("name > ?", name)
	return us.parent
}

func (us *_userQuery[Query]) NameLT(name string) *Query {
	us.IDBExec.Where("name < ?", name)
	return us.parent
}

func (us *_userQuery[Query]) NameLTE(name string) *Query {
	us.IDBExec.Where("name <= ?", name)
	return us.parent
}

func (us *_userQuery[Query]) NameGTE(name string) *Query {
	us.IDBExec.Where("name >= ?", name)
	return us.parent
}

func (us *_userQuery[Query]) NameLike(name string) *Query {
	us.IDBExec.Where("name like ?", "%"+name+"%")
	return us.parent
}

func (us *_userQuery[Query]) NameIn(names []string) *Query {
	us.IDBExec.Where("name in ?", names)
	return us.parent
}

func (us *_userQuery[Query]) NameNEq(name string) *Query {
	us.IDBExec.Where("name != ?", name)
	return us.parent
}

func (us *_userQuery[Query]) NameNotIn(names []string) *Query {
	us.IDBExec.Where("name not in ?", names)
	return us.parent
}
