package main

import (
	"early-setup/config"
	"early-setup/routes"
	"os"
)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	conf := config.Config
	db := config.ConnectDatabase(conf.Database)
	defer config.CloseDatabaseConnection(db)
	// validate := validator.New()

	routes.
		NewRouter(conf.GinMode).
		Run(conf.ServiceHost + ":" + conf.ServicePort)
}