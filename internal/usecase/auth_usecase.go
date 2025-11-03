package usecase

import (
	"fmt"
	"naro-app-be/internal/model"
	"naro-app-be/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Login(email, password string) (*string, error)
	Register(dto *model.CreateUserRequestDto) error
}

type AuthUsecaseImpl struct {
	repo repository.UserRepository
}

func NewAuthUsecase(repo repository.UserRepository) *AuthUsecaseImpl {
	return &AuthUsecaseImpl{repo: repo}
}

func (u *AuthUsecaseImpl) Register(dto *model.CreateUserRequestDto) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)

	if err != nil {
		return fmt.Errorf(`failed to register password %v`, err)
	}

	userMapping := model.CreateUserRequestDto{
		Username: dto.Username,
		Password: string(hashedPassword),
		Name:     dto.Name,
	}

	errCreate := u.repo.CreateUser(userMapping)

	if errCreate != nil {
		return fmt.Errorf(`failed to create user`)
	}

	return nil
}

func (u AuthUsecaseImpl) Login(username, password string) (*string, error) {
	user, err := u.repo.FindUserByUsername(username)

	if err != nil {
		return nil, err
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))

	if errPassword != nil {
		return nil, fmt.Errorf(`mismatch username or password %v`, err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString("maslomanjwt")

	if err != nil {
		return nil, err
	}

	return &tokenString, nil

}
