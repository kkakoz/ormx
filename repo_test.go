package ormx_test

import (
	"context"
	"fmt"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/spf13/viper"
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

func (UserRepo) GetUser(id uint) (*User, error) {
	db := ormx.DB(context.TODO())
	var res *User
	err := db.Where("id = ?", 1500).Find(&res).Error
	return res, err
}

var _ IUserRepo = (*UserRepo)(nil)

func TestRepo(t *testing.T) {
	viper.SetConfigFile("configs/conf.yaml")
	viper.ReadInConfig()
	ormx.New(viper.GetViper())
	userRepo := &UserRepo{ormx.NewRepo[User]()}
	user, err := userRepo.GetUser(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	fmt.Println(userRepo.PageList(context.TODO(), 2, 5, opt.Like("name", "张")))
	fmt.Println(userRepo.GetById(context.TODO()))
}
