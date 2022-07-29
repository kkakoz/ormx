package ormx_test

import (
	"github.com/kkakoz/ormx"
	"github.com/spf13/viper"
	"testing"
)

type User struct {
	ID   uint
	Name string
}

type UserRepo struct {
	ormx.IRepo[User]
}

func TestMain(m *testing.M) {
	viper.SetConfigFile("configs/conf.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	_, err = ormx.New(viper.GetViper())
	if err != nil {
		panic(err)
	}
	ormx.FlushDB()
	m.Run()
}
