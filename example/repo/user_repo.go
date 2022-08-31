package repo

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/example/model"
)

type UserRepo struct {
}

func (u *UserRepo) Query(ctx context.Context) *userQuery {
	return NewUserQuery(ctx)
}

func (u *UserRepo) Update(ctx context.Context) *userUpdate {
	return NewUserUpdate(ctx)
}

func (u *UserRepo) Create(ctx context.Context) *ormx.DBXCreate[model.User] {
	return NewUserCreate(ctx)
}

func (u *UserRepo) Delete(ctx context.Context) *userDelete {
	return NewUserDelete(ctx)
}
