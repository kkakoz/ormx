package ormx

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

// 可以看 https://github.com/win5do/go-microservice-demo/blob/main/docs/sections/gorm.md
type ctxTransactionKey struct{}

type transactionOptions struct {
	sqlOption *sql.TxOptions // 数据库事务
	withNewTx bool           // 在事务内开启一个新的事务
}

type transactionOption func(opt *transactionOptions)

func WithSqlOption(options *sql.TxOptions) transactionOption {
	return func(opt *transactionOptions) {
		opt.sqlOption = options
	}
}

func WithNewTx() transactionOption {
	return func(opt *transactionOptions) {
		opt.withNewTx = true
	}
}

func CtxWithTransaction(ctx context.Context, tx *gorm.DB) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxTransactionKey{}, tx)
}

func Transaction(ctx context.Context, fc func(txctx context.Context) error, options ...transactionOption) error {

	opt := &transactionOptions{}
	for _, o := range options {
		o(opt)
	}

	if !opt.withNewTx { // 开启新事务
		existTx := getTx(ctx)
		if existTx != nil { // 已经在事务里面
			return fc(ctx)
		}
	}

	if opt.sqlOption == nil {
		return ormDB.Transaction(func(tx *gorm.DB) error {
			txctx := CtxWithTransaction(ctx, tx)
			return fc(txctx)
		})
	} else {
		return ormDB.Transaction(func(tx *gorm.DB) error {
			txctx := CtxWithTransaction(ctx, tx)
			return fc(txctx)
		}, opt.sqlOption)
	}

}

func Begin(ctx context.Context, options ...transactionOption) (context.Context, CheckError) {

	opt := &transactionOptions{}
	for _, o := range options {
		o(opt)
	}

	if !opt.withNewTx { // 不默认开启新事务
		existTx := getTx(ctx)
		if existTx != nil {
			return ctx, func(err error) error {
				return err
			}
		}
	}

	var tx *gorm.DB
	if opt.sqlOption != nil {
		tx = ormDB.Begin(opt.sqlOption)
	} else {
		tx = ormDB.Begin()
	}

	return context.WithValue(ctx, ctxTransactionKey{}, tx), func(err error) error {
		if err != nil {
			tx.Rollback()
			return err
		}
		return tx.Commit().Error
	}
}

func DB(ctx context.Context) *gorm.DB {
	iface := ctx.Value(ctxTransactionKey{})

	if iface != nil {
		tx, ok := iface.(*gorm.DB)
		if ok {
			return tx
		}
	}

	return ormDB.WithContext(ctx)
}

func getTx(ctx context.Context) *gorm.DB {
	iface := ctx.Value(ctxTransactionKey{})

	if iface != nil {
		tx, ok := iface.(*gorm.DB)
		if ok {
			return tx
		}
	}
	return nil
}
