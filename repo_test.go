package ormx

import (
	"context"
	"fmt"
	"github.com/kkakoz/ormx/opts"
	"testing"
)

type User struct {
	ID   uint
	Name string
}

type IUserRepo interface {
	IRepo[User]
}

type UserRepo struct {
	repo[User]
}

var _ IUserRepo = (*UserRepo)(nil)

func TestRepo(t *testing.T) {
	userRepo := UserRepo{}
	userList, _ := userRepo.GetList(context.TODO(), opts.NewOpts().Limit(10).Offset(10).Where("name like ?", "1")...)
	user, _ := userRepo.Get(context.TODO(), opts.Where("id = ?", 1))
	fmt.Println(userList)
	fmt.Println(user)
}
