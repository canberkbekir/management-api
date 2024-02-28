package main

import (
	_ "management-api/docs"
	"management-api/internal/config"
	"management-api/internal/router"
	"management-api/internal/util"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server Petstore server.
//	@termsOfService	http://swagger.io/terms/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath	/v1
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
