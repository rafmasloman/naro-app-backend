package route

import (
	"naro-app-be/internal/delivery/http"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App            *gin.Engine
	AuthController *http.AuthControllerImpl
	UserController *http.UserControllerImpl
}

func (c *RouteConfig) SetupRoute() {
	api := c.App.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/login", c.AuthController.Login)
		auth.POST("/register", c.AuthController.Register)
	}

	users := api.Group("/users")
	{
		users.GET("/", c.UserController.FindAllUsers)
	}
}
