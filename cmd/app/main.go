package main

import (
	"flag"
	"fmt"
	"github.com/buts00/Graph/internal/app"
	"github.com/buts00/Graph/internal/config"
	"github.com/buts00/Graph/internal/database"
	"log"

	_ "github.com/lib/pq"
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
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	// connect to database
	databaseCfg := cfg.Database
	db, err := database.NewPostgresDB(databaseCfg.Host, databaseCfg.Port, databaseCfg.User,
		databaseCfg.Password, databaseCfg.DbName)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = db.DB.Close(); err != nil {
			log.Fatal(err, "problem with connection")
		}
	}()

	// Print values
	
	if err = app.PrintNodes(db); err != nil {
		log.Fatal(err)
	}

	if err = app.PrintEdges(db); err != nil {
		log.Fatal(err)
	}

}
