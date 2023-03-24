package usecase

import "golangbin/internal/repository"

type User interface {
	Fetch()
}

type userUsecase struct {
	userRepository repository.User
}

func NewUserUsecases(userRepository repository.User) User {
	return userUsecase{
		userRepository: userRepository,
	}
}

func (u userUsecase) Fetch() {}
