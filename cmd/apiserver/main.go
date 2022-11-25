package main

import (
	"log"
	"os"

	"github.com/golovpeter/http-testapi/internal/app/apiserver"
	"github.com/joho/godotenv"
)

var (
	configPath string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	config := apiserver.NewConfig()

	bindAddr := os.Getenv("BIND_ADDR")
	if bindAddr != "" {
		config.BindAddr = os.Getenv("BIND_ADDR")
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel != "" {
		config.LogLevel = os.Getenv("LOG_LEVEL")
	}

	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatalln(err)
	}
}
