package repo_test

import (
	"context"
	"fmt"
	"github.com/kkakoz/ormx"
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
	userRepo := repo.UserRepo{}

	ctx := context.TODO()

	err := userRepo.Create(ctx).Add(&model.User{Name: "1"})
	if err != nil {
		panic(err)
	}

	list, err := userRepo.Query(ctx).NameLike("zhangsan").List()
	if err != nil {
		panic(err)
	}
	fmt.Println(list)

	user, err := userRepo.Query(ctx).ID(1).One()
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	err = userRepo.Update(ctx).Update("name", "lisi")
	if err != nil {
		panic(err)
	}

	err = userRepo.Delete(ctx).Delete()
	if err != nil {
		panic(err)
	}

}

func TestInit(t *testing.T) {

	ormx.QueryInit(model.User{}, os.Stdout)

}