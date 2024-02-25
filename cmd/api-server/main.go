package main

import (
	"github.com/labstack/echo/v4"
	"management-api/internal/config"
	"management-api/internal/controller"
	"management-api/internal/repository"
	"management-api/internal/service"
	"management-api/internal/util"
)

func main() {

	util.InitLog()
	serverConfig := config.NewConfig().Server
	coucheBaseConfig := config.NewConfig().Couchbase
	e := echo.New()

	cluster, err := config.NewCouchebaseClient(&coucheBaseConfig)
	if err != nil {
		util.Logger.Fatal().Err(err).Msg("Error connecting to couchbase")
	}

	userRepository := repository.NewUserRepository(cluster)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userController.Register(e)

	util.Logger.Fatal().Err(e.Start(":" + serverConfig.Port)).Msg("Error starting server")

}
