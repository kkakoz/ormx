package ormx

import (
	"context"
	"testing"
)

type User struct {
	ID   uint
	Name string
}

func TestRepo(t *testing.T) {
	userRepo := Repo[User]{}
	userRepo.GetById(context.TODO(), 1)
}
