package model

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserState uint

const (
	UserStateA UserState = 1
)

type Class struct {
	ID uint `json:"id"`

	Users []User `json:"users" ormx:"with:Users;"`
}
