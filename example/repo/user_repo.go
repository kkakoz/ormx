package repo

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/example/model"
)

type userRepo struct {
}

func UserRepo() *userRepo {
	return &userRepo{}
}

func (u *userRepo) Query(ctx context.Context) *ormx.DBXQuery[model.User] {
	return ormx.NewDBXQuery[model.User](ctx)
}

func (u *userRepo) Update(ctx context.Context) *ormx.DBXUpdate[model.User] {
	return ormx.NewDBXUpdate[model.User](ctx)
}

func (u *userRepo) Create(ctx context.Context, users ...*model.User) error {
	dbxCreate := ormx.NewDBXCreate[model.User](ctx)
	return dbxCreate.AddList(users)
}

func (u *userRepo) Delete(ctx context.Context) *ormx.DBXDelete[model.User] {
	return ormx.NewDBXDelete[model.User](ctx)
}
