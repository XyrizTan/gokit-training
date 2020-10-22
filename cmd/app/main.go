package main

import (
	"fmt"
	"os"

	"github.com/viki-org/go-utils/vaultconsul"
)

// Config is the service config, retrieved from consul and vault in memory.
type Config struct {
	Database struct {
		Host               string `consul:"database/host"`
		Port               int    `consul:"database/port"`
		Name               string `consul:"database/name"`
		MaxIdleConnections int    `consul:"database/max_idle_connections"`
		MaxOpenConnections int    `consul:"database/max_open_connections"`
		LogMode            bool   `consul:"database/log_mode"`
		User               string `vault:"database.user"`
		Password           string `vault:"database.password"`
	}

	Server struct {
		Port              int    `consul:"server/port"`
		LoggingLevel      string `consul:"server/logging_level"`
		HoneybadgerAPIKey string `vault:"honeybadger.api_key"`
	}

	Queue struct {
		Host     string `consul:"queue/host"`
		Port     int    `consul:"queue/port"`
		Name     string `consul:"queue/name"`
		Username string `vault:"queue.username"`
		Password string `vault:"queue.password"`
	}

	Redis struct {
		Host               string `consul:"redis/host"`
		Port               int    `consul:"redis/port"`
		DB                 int    `consul:"redis/db"`
		MaxIdleConnections int    `consul:"redis/max_idle_connections"`
		IdleTimeout        int    `consul:"redis/idle_timeout"`
	}
}

func main() {
	var config Config
	err := vaultconsul.Decode(&config)
	if err != nil {
		fmt.Println("Vault&Consul decode", err)
		os.Exit(1)
	}
	app := App{}
	err = app.Initialize(
		WithDB(config),
		WithQueue(config),
		WithRedis(config),
	)
	if err != nil {
		fmt.Println("Server Initialisation", err)
		os.Exit(1)
	}
	app.Run(config)
}
