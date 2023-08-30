package main

import (
	_ "github.com/burravlev/avito-tech-test/docs"

	"github.com/burravlev/avito-tech-test/cmd/app"
)

var configPath = "config.yml"

//	@title			Avito A/B testing service
//	@version		1.0
//	@desccription	Microservice for sement users

//	@host	localhost:8080
//
// BasePath /
func main() {
	app.Run(configPath)
}
