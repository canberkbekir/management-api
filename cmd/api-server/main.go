package main

import (
	"management-api/internal/config"
	"management-api/internal/repository"
	"management-api/internal/util"
	"net/http"
	"time"
)

func main() {
	util.InitLog()
	serverConfig := config.NewConfig().Server
	server := &http.Server{
		Addr:           ":" + serverConfig.Port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	coucheBaseConfig := config.NewConfig().Couchbase
	test, err := config.NewCouchebaseClient(&coucheBaseConfig)
	if err != nil {
		util.Logger.Fatal().Err(err).Msg("Error connecting to couchbase")
	}

	repository.NewUserRepository(test)
	result, _ := repository.NewUserRepository(test).GetAll()
	util.Logger.Info().Msgf("Result: %v", result)
	util.Logger.Info().Msgf("Server started on port %s", serverConfig.Port)
	if err := server.ListenAndServe(); err != nil {
		util.Logger.Fatal().Err(err).Msg("Server failed to start")
	}
}
