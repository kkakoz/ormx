package ormx

type ErrHandle func(err error) error

var DefaultErrHandle = func(err error) error {
	return err
}
