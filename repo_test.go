package ormx_test

import (
	"context"
	"fmt"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opts"
	"testing"
)

type User struct {
	ID   uint
	Name string
}

type IUserRepo interface {
	ormx.IRepo[User]
}

type UserRepo struct {
	ormx.IRepo[User]
}

var _ IUserRepo = (*UserRepo)(nil)

func TestRepo(t *testing.T) {
	var userRepo IUserRepo = ormx.NewRepo[User]()
	userList, _ := userRepo.GetList(context.TODO(), opts.NewOpts().Limit(10).Offset(10).Where("name like ?", "1")...)
	user, _ := userRepo.Get(context.TODO(), opts.Where("id = ?", 1))
	fmt.Println(userList)
	fmt.Println(user)
}
