package main

import (
	"flag"

	"github.com/tekpriest/poprev/cmd/config"
	"github.com/tekpriest/poprev/cmd/database"
)

func main() {
	env := flag.String("env", "dev", "environment to run on")
	flag.Parse()

	config := config.LoadConfig(*env)
	database.NewConnection(config)
}
