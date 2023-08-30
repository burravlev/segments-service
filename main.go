package main

import (
	_ "github.com/burravlev/avito-tech-test/docs"

	"github.com/burravlev/avito-tech-test/cmd/app"
)

var configPath = "config.yml"

func main() {
	app.Run(configPath)
}
