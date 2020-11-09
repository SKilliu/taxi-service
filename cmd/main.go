package main

import (
	"fmt"

	app "github.com/SKilliu/taxi-service"

	"github.com/SKilliu/taxi-service/config"
	"github.com/SKilliu/taxi-service/docs"
	"github.com/pkg/errors"
)

const pathToConfigFile = "./static/envs.yaml"

// @title Taxi-service
// @version 1.0
// @securityDefinitions.apiKey bearerAuth
// @in header
// @name Authorization
func main() {
	apiConfig := config.New()
	log := apiConfig.Log()

	config.UploadEnvironmentVariables(pathToConfigFile)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", apiConfig.HTTP().Host, apiConfig.HTTP().Port)

	api := app.New(apiConfig)
	if err := api.Start(); err != nil {
		log.WithError(err)
		panic(errors.Wrap(err, "failed to start api"))
	}
}
