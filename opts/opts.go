package opts

import (
	"fmt"

	"gorm.io/gorm"
)

type Option func(db *gorm.DB) *gorm.DB

type Options []Option

func NewOpts() Options {
	return Options{}
}

func (o Options) Where(key string, value any) Options {
	o = append(o, Where(key, value))
	return o
}

func Where(key string, values ...any) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(key, values...)
	}
}

func (o Options) Like(key string, value string) Options {
	o = append(o, Like(key, value))
	return o
}

func (o Options) Func(f Option) Options {
	o = append(o, f)
	return o
}

func Like(key string, value string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s like ?", key), "%"+value+"%")
	}
}

func (o Options) IsLike(is bool, key string, value string) Options {
	o = append(o, IsLike(is, key, value))
	return o
}

func IsLike(is bool, key string, value string) Option {
	return func(db *gorm.DB) *gorm.DB {
		if is {
			return db.Where(fmt.Sprintf("%s like ?", key), "%"+value+"%")
		}
		return db
	}
}

func (o Options) EQ(key string, value any) Options {
	o = append(o, EQ(key, value))
	return o
}

func EQ(key string, value any) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s = ?", key), value)
	}
}

func (o Options) IsEQ(is bool, key string, value any) Options {
	o = append(o, IsEQ(is, key, value))
	return o
}

func IsEQ(is bool, key string, value any) Option {
	return func(db *gorm.DB) *gorm.DB {
		if is {
			return db.Where(fmt.Sprintf("%s = ?", key), value)
		}
		return db
	}
}

func (o Options) In(key string, value any) Options {
	o = append(o, In(key, value))
	return o
}

func In(key string, value any) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s in ?", key), value)
	}
}

func (o Options) Page(page int, pageSize int) Options {
	o = append(o, Page(page, pageSize))
	return o
}

func Page(page int, pageSize int) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(pageSize).Offset((page - 1) * pageSize)
	}
}

func (o Options) Order(key string) Options {
	o = append(o, Order(key))
	return o
}

func Order(key string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(key)
	}
}
