package repository

import "context"

type User interface {
	Get(context.Context)
}

type userRepository struct {
}

func NewUserRepository() User {
	return userRepository{}
}
func (r userRepository) Get(context.Context) {

}
