package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	_ "management-api/docs"
	"management-api/internal/config"
	"management-api/internal/router"
	"management-api/internal/util"
	"os"
)

//	@title			Management API
//	@version		1.0
//	@description	Testing Swagger APIs.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@securityDefinitions.apiKey	JWT
//	@in							header
//	@name						token

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/api
func main() {

	util.InitLog()
	serverConfig := config.NewConfig().Server
	coucheBaseConfig := config.NewConfig().Couchbase

	cluster, err := config.NewCouchebaseClient(&coucheBaseConfig)
	if err != nil {
		util.Logger.Fatal().Err(err)
	}

	e := router.Init(cluster)
	util.Logger.Fatal().Err(e.Start(":" + serverConfig.Port)).Msg("Error starting server")

}

func writeRoutesFile(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		util.Logger.Fatal().Err(err)
	}

	err = os.WriteFile("routes.json", data, 0777)
	if err != nil {
		util.Logger.Fatal().Err(err)
	}
}
