package ormx_test

import (
	"github.com/kkakoz/ormx"
	"github.com/spf13/viper"
	"testing"
)

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

func TestB(t *testing.T) {
	//userQuery := ormx.NewQuery[UserQuery, User](UserQuery{}, ormx.DB(context.Background()))
}

type User struct {
}

type UserQuery struct {
	ormx.Query
}

func (q *UserQuery) IDEQ() {
	q.Query.Where("id = ?")
}
