package main

import (
	"management-api/internal/config"
	"management-api/internal/router"
	"management-api/internal/util"
)

func main() {

	util.InitLog()
	serverConfig := config.NewConfig().Server
	coucheBaseConfig := config.NewConfig().Couchbase

	cluster, err := config.NewCouchebaseClient(&coucheBaseConfig)
	if err != nil {
		util.Logger.Fatal().Err(err).Msg("Error connecting to couchbase")
	}

	e := router.Init(cluster)
	util.Logger.Fatal().Err(e.Start(":" + serverConfig.Port)).Msg("Error starting server")

}
