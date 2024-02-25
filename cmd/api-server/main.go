package main

import (
	"management-api/internal/config"
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
	_, err := config.NewCouchebaseClient(&coucheBaseConfig)
	if err != nil {
		util.Logger.Fatal().Err(err).Msg("Error connecting to couchbase")
	}

	util.Logger.Info().Msgf("Server started on port %s", serverConfig.Port)
	if err := server.ListenAndServe(); err != nil {
		util.Logger.Fatal().Err(err).Msg("Server failed to start")
	}
}
