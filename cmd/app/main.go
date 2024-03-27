package main

import (
	"flag"
	"fmt"
	"github.com/buts00/Graph/internal/app/apiserver"
	"github.com/buts00/Graph/internal/database"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs", "path to config")
}

func main() {

	//set Config

	if err := initConfig(); err != nil {
		log.Fatal("error loading config: " + err.Error())
	}

	if err := gotenv.Load(); err != nil {
		log.Fatal("cannot load env variables: " + err.Error())
	}

	//connect to db
	db, err := database.NewPostgresDB(viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.user"),
		os.Getenv("DB_PASSWORD"), viper.GetString("db.db_name"))

	if err != nil {
		log.Fatal("cannot connect to database: " + err.Error())
	}

	defer func() {
		if err = db.DB.Close(); err != nil {
			log.Fatal("problem with closing db: " + err.Error())
		}
	}()

	// Start Server
	bindAddr := viper.GetString("bind_addr")
	fmt.Println("Server run on port ", bindAddr)
	if err := apiserver.Start(bindAddr, db); err != nil {
		log.Fatal("error occurred while running http server" + err.Error())
	}
}

func initConfig() error {
	flag.Parse()
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
