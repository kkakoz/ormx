package ormx

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var ormDB *gorm.DB

func InitDB(database *gorm.DB) {
	ormDB = database
}

func New(viper *viper.Viper) (*gorm.DB, error) {
	var err error
	viper.SetDefault("db.user", "root")
	viper.SetDefault("db.password", "")
	viper.SetDefault("db.host", "127.0.0.1")
	viper.SetDefault("db.port", "3306")
	viper.SetDefault("db.name", "test")
	viper.SetDefault("db.console_log", "false")
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?"+
		"charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.user"), viper.GetString("db.password"),
		viper.GetString("db.host"), viper.GetString("db.port"),
		viper.GetString("db.name"))
	loggerLevel := logger.Warn
	if viper.GetBool("db.log_info") { // 控制台打印普通sql
		loggerLevel = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  loggerLevel, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	config := &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true, // AutoMigrate不会自动添加外键
	}

	ormDB, err = gorm.Open(mysql.Open(dns), config)
	return ormDB, err
}

// 可以看 https://github.com/win5do/go-microservice-demo/blob/main/docs/sections/gorm.md
type ctxTransactionKey struct{}

func CtxWithTransaction(ctx context.Context, tx *gorm.DB) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxTransactionKey{}, tx)
}

func Transaction(ctx context.Context, fc func(txctx context.Context) error) error {
	db := ormDB.WithContext(ctx)

	return db.Transaction(func(tx *gorm.DB) error {
		txctx := CtxWithTransaction(ctx, tx)
		return fc(txctx)
	})
}

func Begin(ctx context.Context, opts ...*sql.TxOptions) (context.Context, CheckError) {
	tx := ormDB.Begin(opts...)
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

func CheckErr(database *gorm.DB, err error) error {
	if err != nil {
		curErr := database.Rollback().Error
		if curErr != nil {
			return curErr
		}
		return err
	}
	return database.Commit().Error
}

type Model struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FlushDB() {

	var tables []string
	err := ormDB.Table("information_schema.tables").Where("table_schema = ?", "server_test").Pluck("table_name", &tables).Error
	if err != nil {
		log.Fatalln(err)
	}

	for _, table := range tables {
		ormDB.Table(table).Exec(fmt.Sprintf("truncate table `%s`", table))
	}

}

type CheckError func(err error) error
