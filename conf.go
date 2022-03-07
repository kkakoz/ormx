package ormx

type ErrHandler func(err error) error

var DefaultErrHandler = func(err error) error {
	return err
}
