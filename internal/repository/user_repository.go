package repository

import (
	"fmt"
	"naro-app-be/internal/entity"
	"naro-app-be/internal/model"
	"naro-app-be/internal/model/mapping"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(dto model.CreateUserRequestDto) error
	FindUserByUsername(username string) (*model.UserResponse, error)
	FindAllUsers() (*[]model.UserResponse, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) CreateUser(dto model.CreateUserRequestDto) error {

	userMapping := mapping.UserMappingRequestDto(&dto)

	model := r.db.Create(userMapping)

	return model.Error

}

func (r *UserRepositoryImpl) FindUserByUsername(username string) (*model.UserResponse, error) {
	var user entity.User

	if err := r.db.Where(`username = ?`, username).Find(&user).Error; err != nil {
		return nil, fmt.Errorf(`username not found %v`, err)
	}

	userMapping := mapping.UserMappingResponseDto(&user)

	return userMapping, nil
}

func (r *UserRepositoryImpl) FindAllUsers() (*[]model.UserResponse, error) {
	var user []entity.User

	if err := r.db.Find(&user).Error; err != nil {
		return nil, fmt.Errorf(`users not found %v`, err)
	}

	mappingUser := mapping.UsersMappingResponseDTO(user)

	return mappingUser, nil

}
