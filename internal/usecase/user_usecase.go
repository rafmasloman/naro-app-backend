package usecase

import (
	"naro-app-be/internal/model"
	"naro-app-be/internal/repository"
)

type UserUsecase interface {
	FindAllUsers() (*[]model.UserResponse, error)
}

type UserUsecaseImpl struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{repo: repo}
}

func (u *UserUsecaseImpl) FindAllUsers() (*[]model.UserResponse, error) {
	return u.repo.FindAllUsers()
}
