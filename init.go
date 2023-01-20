package ormx

import (
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

type option struct {
	log logger.Interface
}

type optionFunc func(option *option) *option

func WithLogger(log logger.Interface) optionFunc {
	return func(opt *option) *option {
		opt.log = log
		return opt
	}
}

func InitDB(database *gorm.DB) {
	ormDB = database
}

func New(viper *viper.Viper, opts ...optionFunc) (*gorm.DB, error) {
	opt := &option{
		log: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Warn, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,       // Disable color
			},
		),
	}
	for _, optFunc := range opts {
		opt = optFunc(opt)
	}
	var err error
	viper.SetDefault("db.user", "root")
	viper.SetDefault("db.password", "")
	viper.SetDefault("db.host", "127.0.0.1")
	viper.SetDefault("db.port", "3306")
	viper.SetDefault("db.name", "test")
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?"+
		"charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.user"), viper.GetString("db.password"),
		viper.GetString("db.host"), viper.GetString("db.port"),
		viper.GetString("db.name"))

	config := &gorm.Config{
		Logger:                                   opt.log,
		DisableForeignKeyConstraintWhenMigrating: true, // AutoMigrate不会自动添加外键
	}

	ormDB, err = gorm.Open(mysql.Open(dns), config)
	return ormDB, err
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
