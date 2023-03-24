package engine

import (
	"golangbin/internal/handler"
	"golangbin/internal/repository"
	"golangbin/internal/usecase"
)

type Handler struct {
	User handler.User
}

type Usecase struct {
	User usecase.User
}

type Repository struct {
	User repository.User
}

func initRepository() Repository {
	return Repository{
		User: repository.NewUserRepository(),
	}
}

func initUsecase(r Repository) Usecase {
	return Usecase{
		User: usecase.NewUserUsecases(r.User),
	}
}

func initHandler(r Usecase) Handler {
	return Handler{
		User: handler.NewUserHandler(),
	}
}
