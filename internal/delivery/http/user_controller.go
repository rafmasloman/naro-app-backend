package http

import (
	"naro-app-be/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	usecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) *UserControllerImpl {
	return &UserControllerImpl{
		usecase: usecase,
	}
}

func (ctl *UserControllerImpl) FindAllUsers(c *gin.Context) {
	users, err := ctl.usecase.FindAllUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"data":       users,
		"message":    "login success",
	})
}
