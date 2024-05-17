package main

import (
	"basic-api/internal/app"
	maria "basic-api/internal/core/database/mariadb"
	"basic-api/internal/core/logger"
	"basic-api/internal/core/server"
	"basic-api/internal/env"
	"basic-api/internal/router"

	"github.com/sirupsen/logrus"
)

func main() {
	envConfig, err := env.Read("configs")
	if err != nil {
		panic(err)
	}

	ds, err := maria.NewMariaDB(&maria.Config{
		Host:         envConfig.Config.Database.Host,
		Port:         envConfig.Config.Database.Port,
		User:         envConfig.Config.Database.Username,
		Password:     envConfig.Config.Database.Password,
		DatabaseName: envConfig.Config.Database.Name,
		Debug:        !envConfig.Config.Release,
	})
	if err != nil {
		panic(err)
	}

	// Mongodb connection .. not implement yet
	// err = mongo.InitDatabase(&mongo.Options{
	// 	URL:          envConfig.DatabaseURL,
	// 	Port:         envConfig.DatabasePort,
	// 	DatabaseName: envConfig.DatabaseName,
	// 	Username:     envConfig.DatabaseUsername,
	// 	Password:     envConfig.DatabasePassword,
	// 	Root:         envConfig.DatabaseRoot,
	// 	Debug:        !envConfig.Release,
	// 	IsProd:       envConfig.Production,
	// })
	// if err != nil {
	// 	panic(err)
	// }

	options := &router.Options{
		AppContext: &app.Context{
			Db:     ds.Db,
			Config: envConfig.Config,
			// MongoDb: mds,
		},
	}
	logrus.SetOutput(&logger.OutputSplitter{})

	server.New(router.NewWithOptions(options), envConfig.Config.ServerPort).Start()
}
