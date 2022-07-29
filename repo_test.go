package ormx_test

import (
	"context"
	"fmt"
	"github.com/kkakoz/ormx"
	"testing"
)

func TestRepo(t *testing.T) {
	userRepo := &UserRepo{ormx.NewRepo[User]()}
	user, err := userRepo.GetById(context.TODO(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
