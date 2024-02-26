package router

import (
	"github.com/couchbase/gocb/v2"
	"github.com/labstack/echo/v4"
	"management-api/internal/controller"
	"management-api/internal/middleware"
	"management-api/internal/repository"
	"management-api/internal/service"
)

func Init(cluster *gocb.Cluster) *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggingMiddleware)

	// Initialize the controller, service, and repository for the user
	userRepository := repository.NewUserRepository(cluster)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	api := e.Group("/api")
	{
		api.Group("/user")
		{
			api.GET("", userController.GetAllUser)
			api.GET("/id/:id", userController.GetUserById)
			api.POST("", userController.UpsertUser)
			api.DELETE("/id/:id", userController.DeleteUser)
		}
	}

	return e
}
