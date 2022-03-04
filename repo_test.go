package ormx

import (
	"fmt"
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
	Repo[User]
}

var _ IUserRepo = (*UserRepo)(nil)

func TestRepo(t *testing.T) {
	fmt.Println("123")
}
