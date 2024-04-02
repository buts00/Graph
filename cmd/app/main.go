package main

import (
	"flag"

	"github.com/buts00/Graph/internal/app/apiserver"
	config2 "github.com/buts00/Graph/internal/config"
	"github.com/buts00/Graph/internal/database"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
	"os"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs", "path to config")
}

func main() {
	// Initialize logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	// Set Config
	if err := gotenv.Load(); err != nil {
		logrus.Fatal("cannot load env variables: ", err)
	}
	flag.Parse()
	config, err := config2.LoadConfig(configPath)
	if err != nil {
		logrus.Fatal("cannot set up config: ", err)

	}

	// Connect to db
	db, err := database.NewPostgresDB(config.Database.Host, config.Database.Port, config.Database.User,
		config.Database.Password, config.Database.DbName)
	if err != nil {
		logrus.Fatal("cannot connect to database: ", err)
	}

	defer func() {
		if err = db.DB.Close(); err != nil {
			logrus.Fatal("problem with closing db: ", err)
		}
	}()

	// Start Server

	logrus.Info("Server run on port ", config.Server.BindAddr)
	if err := apiserver.Start(config.Server.BindAddr, db); err != nil {
		logrus.Fatal("error occurred while running http server: ", err)

	}
}
