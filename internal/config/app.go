package config

import (
	"naro-app-be/internal/delivery/http"
	"naro-app-be/internal/repository"
	"naro-app-be/internal/usecase"

	"gorm.io/gorm"
)

type AppResponse struct {
	AuthController http.AuthControllerImpl
	UserController http.UserControllerImpl
}

func Bootstrap(db *gorm.DB) *AppResponse {

	userRepository := repository.NewUserRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepository)
	userUsecase := usecase.NewUserUsecase(userRepository)

	authController := http.NewAuthController(authUsecase)
	userController := http.NewUserController(userUsecase)

	return &AppResponse{
		AuthController: *authController,
		UserController: *userController,
	}

}
