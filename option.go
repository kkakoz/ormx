package ormx

type Option[T any] func(r *repo[T])

func ErrOption[T any](f func(error) error) Option[T] {
	return func(r *repo[T]) {
		r.errHandle = f
	}
}
