package ormx_test

import (
	"context"
	"errors"
	"github.com/kkakoz/ormx"
	"log"
	"testing"
)

func TestTransaction(t *testing.T) {
	errMsg := "未知错误"

	userRepo := &UserRepo{ormx.NewRepo[User]()}
	ctx := context.TODO()
	err := ormx.Transaction(ctx, func(ctx context.Context) error {
		err := userRepo.Add(ctx, &User{Name: "张三"})
		if err != nil {
			return err
		}
		err = ormx.Transaction(ctx, func(txctx context.Context) error {
			err := userRepo.Add(ctx, &User{Name: "李四"})
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
		return errors.New(errMsg)
	})

	if err.Error() != errMsg {
		log.Fatal("未能捕获错误")
	}

}

func TestWithNewTx(t *testing.T) {
	errMsg := "未知错误"

	u1 := &User{Name: "张三"}
	u2 := &User{Name: "李四"}

	userRepo := &UserRepo{ormx.NewRepo[User]()}
	ctx := context.TODO()
	err := ormx.Transaction(ctx, func(ctx context.Context) error {
		err := userRepo.Add(ctx, u1)
		if err != nil {
			return err
		}
		err = ormx.Transaction(ctx, func(txctx context.Context) error {
			err := userRepo.Add(txctx, u2) // u2 保存 不rollback
			if err != nil {
				return err
			}
			return nil
		}, ormx.WithNewTx())
		if err != nil {
			return err
		}
		return errors.New(errMsg) // u1 返回 err rollback
	})

	if err.Error() != errMsg {
		log.Fatal("未能捕获错误")
	}

}
