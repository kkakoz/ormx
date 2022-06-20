package opts_test

import (
	"context"
	"fmt"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opts"
	"log"
	"testing"
)

type User struct {
	ID    uint
	Name  string
	Phone string
}

func GetList(opts ...opts.Option) ([]*User, error) {
	db := ormx.DB(context.TODO())
	for _, opt := range opts {
		opt(db)
	}
	users := []*User{}
	err := db.Find(&users).Error
	return users, err
}

func TestGetList(t *testing.T) {
	name := "张"
	list, err := GetList(opts.Like("name", name))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(list)
	phone := "12345678910"
	list, err = GetList(opts.Where("phone = ?", phone))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(list)
}
