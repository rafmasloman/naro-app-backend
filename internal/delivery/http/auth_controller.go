package http

import (
	"naro-app-be/internal/model"
	"naro-app-be/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
}

type AuthControllerImpl struct {
	usecase usecase.AuthUsecase
}

func NewAuthController(usecase usecase.AuthUsecase) *AuthControllerImpl {
	return &AuthControllerImpl{
		usecase: usecase,
	}
}

func (ctl *AuthControllerImpl) Login(c *gin.Context) {
	var req model.LoginAuthRequestDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctl.usecase.Login(req.Username, req.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
			"message":    "Username atau Password salah",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"message": "login success",
	})
}

func (ctl *AuthControllerImpl) Register(c *gin.Context) {
	var req model.CreateUserRequestDto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctl.usecase.Register(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
			"message":    "Username atau Password salah",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Anda berhasil mendaftar",
	})
}
