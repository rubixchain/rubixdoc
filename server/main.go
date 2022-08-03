package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/EnsurityTechnologies/apiconfig"
	"github.com/EnsurityTechnologies/logger"
)

func main() {
	// cfg := Config{
	// 	Config: config.Config{
	// 		LogFile:     "log.txt",
	// 		Production:  "false",
	// 		HostAddress: "localhost",
	// 		HostPort:    "12345",
	// 	},
	// }

	var cfg Config
	err := apiconfig.LoadAPIConfig("api_config.json", "TestKey@Prod#$erver^2021", &cfg)
	if err != nil {
		panic(err)
	}

	fp, err := os.OpenFile(cfg.LogFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	logOptions := &logger.LoggerOptions{
		Name:   "Main",
		Color:  logger.AutoColor,
		Output: fp,
	}

	log := logger.New(logOptions)
	s, err := NewServer(&cfg, log)
	if err != nil {
		log.Error("Failed to create server")
		return
	}
	log.Info("Starting server...")
	go s.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGINT)

	<-c
	s.Shutdown()
	log.Info("Shutting down...")
}
