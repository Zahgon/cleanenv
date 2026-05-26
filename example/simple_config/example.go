package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config is a application configuration structure
type Config struct {
	Database struct {
		Host        string `yaml:"host" env:"DB_HOST" env-description:"Database host"`
		Port        string `yaml:"port" env:"DB_PORT" env-description:"Database port"`
		Username    string `yaml:"username" env:"DB_USER" env-description:"Database user name"`
		Password    string `env:"DB_PASSWORD" env-description:"Database user password"`
		Name        string `yaml:"db-name" env:"DB_NAME" env-description:"Database name"`
		Connections int    `yaml:"connections" env:"DB_CONNECTIONS" env-description:"Total number of database connections"`
	} `yaml:"database"`
	Server struct {
		Host string `yaml:"host" env:"SRV_HOST,HOST" env-description:"Server host" env-default:"localhost"`
		Port string `yaml:"port" env:"SRV_PORT,PORT" env-description:"Server port" env-default:"8080"`
	} `yaml:"server"`
	Greeting string `env:"GREETING" env-description:"Greeting phrase" env-default:"Hello!"`
}

// Args command-line parameters
type Args struct {
	ConfigPath string
}

// ConnectDB connects to an abstract database
func ConnectDB(host, port, user, password, name string, conn int) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	var cfg Config

	args := ProcessArgs(&cfg)

	// read configuration from the file and environment variables
	if err := cleanenv.ReadConfig(args.ConfigPath, &cfg); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// connect to the DB (example)
	ConnectDB(cfg.Database.Host, cfg.Database.Port, cfg.Database.Username,
		cfg.Database.Password, cfg.Database.Name, cfg.Database.Connections)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", cfg.Greeting)
	})

	http.ListenAndServe(cfg.Server.Host+":"+cfg.Server.Port, nil)
}

// ProcessArgs processes and handles CLI arguments
func ProcessArgs(cfg interface{}) Args { _ = "STUB: not implemented"; return *new(Args) }
