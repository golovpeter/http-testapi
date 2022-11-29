package apiserver

import "github.com/golovpeter/http-testapi/internal/store"

type Config struct {
	BindAddr string
	LogLevel string
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}

func (s *Config) Start() error {
	return nil
}
