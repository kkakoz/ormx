package repo_test

import (
	"context"
	"fmt"
	"github.com/kkakoz/ormx/example/model"
	"github.com/kkakoz/ormx/example/repo"
	"os"
	"testing"
)

type User struct {
	ID        uint   `json:"id"`
	ClassName string `json:"class_name"`
}

func TestUserRepo(t *testing.T) {
	userRepo := repo.UserRepo()

	ctx := context.Background()

	userRepo.Create(ctx, &model.User{ID: 1, Name: "John"})
	one, err := userRepo.Query(ctx).Like("name", "zhangsan").One()
	if err != nil {
		t.Fail()
	}
	fmt.Println(one)

	one, err = userRepo.Query(ctx).EQ("id", "1").One()
	if err != nil {
		t.Fail()
	}

	//var count int64
	//err = userRepo.Delete(ctx).Where("name = ?", "John").Exec(ormx.RowEffect(&count))
	//if err != nil {
	//	t.Fail()
	//}
	//fmt.Println(count)
}

func TestInit(t *testing.T) {

	file, err := os.Open("./to/")
	fmt.Println(file, err)
}
