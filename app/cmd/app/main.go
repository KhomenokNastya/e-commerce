package main

import (
	"golang/internal/app"
	"golang/internal/config"
	"golang/pkg/logging"
	"log"
)

func main() {
	log.Println("config initializing")
	cfg := config.GetConfig()

	log.Println("logger initializing")
	logging.Init(cfg.AppConfig.LogLevel)
	logger := logging.GetLogger()

	a, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("Running Application")
	a.Run()
}
