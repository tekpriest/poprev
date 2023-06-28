package main

import (
	"flag"

	"github.com/tekpriest/poprev/cmd/config"
	"github.com/tekpriest/poprev/cmd/database"
)

//	@title			Poprev API
//	@version		1.0
//	@description	Poprev API

//	@license.name	MIT

//	@securityDefinitions.apiKey	ApiKey
//	@in							header
//	@name						x-auth-token

// @Accept		json
// @Produce	json
// @BasePath	/
func main() {
	env := flag.String("env", "dev", "environment to run on")
	flag.Parse()

	// load config
	c := config.LoadConfig(*env)

	// connect database
	database.NewConnection(c)

	// start server
	NewRouter(c).Serve()
}
