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
	return append(o, Where(key, value))
}

func Where(key string, values ...any) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(key, values...)
	}
}

func (o Options) IsWhere(is bool, key string, values ...any) Options {
	return append(o, IsWhere(is, key, values...))
}

func IsWhere(is bool, key string, values ...any) Option {
	return func(db *gorm.DB) *gorm.DB {
		if is {
			return db.Where(key, values...)
		}
		return db
	}
}

func (o Options) Like(key string, value string) Options {
	return append(o, Like(key, value))
}

func (o Options) Func(f Option) Options {
	return append(o, f)
}

func Like(key string, value string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s like ?", key), "%"+value+"%")
	}
}

func (o Options) IsLike(is bool, key string, value string) Options {
	return append(o, IsLike(is, key, value))
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
	return append(o, EQ(key, value))
}

func EQ(key string, value any) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s = ?", key), value)
	}
}

func (o Options) IsEQ(is bool, key string, value any) Options {
	return append(o, IsEQ(is, key, value))
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
	return append(o, In(key, value))
}

func In(key string, value any) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s in ?", key), value)
	}
}

func (o Options) Limit(limit int) Options {
	return append(o, Limit(limit))
}

func Limit(limit int) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	}
}

func (o Options) Offset(offset int) Options {
	return append(o, Offset(offset))
}

func Offset(offset int) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset)
	}
}

func (o Options) Order(key string) Options {
	return append(o, Order(key))
}

func Order(key string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(key)
	}
}

func (o Options) Preload(name string, args ...any) Options {
	return append(o, Preload(name, args...))
}

func Preload(name string, args ...any) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(name, args...)
	}
}
