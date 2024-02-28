package router

import (
	"github.com/couchbase/gocb/v2"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"management-api/internal/controller"
	"management-api/internal/middleware"
	"management-api/internal/repository"
	"management-api/internal/service"
)

func Init(cluster *gocb.Cluster) *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggingMiddleware)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Initialize the controller, service, and repository for the user
	userRepository := repository.NewUserRepository(cluster)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	api := e.Group("/api")
	{
		user := api.Group("/user")
		{
			user.GET("", userController.GetAllUser)
			user.GET("/id/:id", userController.GetUserById)
			user.POST("", userController.UpsertUser)
			user.DELETE("/id/:id", userController.DeleteUser)
		}
	}

	return e
}
