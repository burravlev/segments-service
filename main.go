package main

import "github.com/burravlev/avito-tech-test/cmd/app"

var configPath = "config.yml"

func main() {
	app.Run(configPath)
}
