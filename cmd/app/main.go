package main

import (
	"flag"
	"github.com/buts00/Graph/internal/app/apiserver"
	"github.com/buts00/Graph/internal/config"
	"github.com/buts00/Graph/internal/database"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config")
}

func main() {

	//set Config
	flag.Parse()
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err, "Error loading config: ")
	}

	//connect to database
	databaseCfg := cfg.Database
	password := os.Getenv("PASSWORD_graph_db")
	db, err := database.NewPostgresDB(databaseCfg.Host, databaseCfg.Port, databaseCfg.User,
		password, databaseCfg.DbName)

	if err != nil {
		log.Fatal(err, "cannot connect to database")
	}

	defer func() {
		if err = db.DB.Close(); err != nil {
			log.Fatal(err, "problem with connection")
		}
	}()

	// Start Server
	if err := apiserver.Start(cfg.Server.BindAddr, db); err != nil {
		log.Fatal(err)
	}

}
