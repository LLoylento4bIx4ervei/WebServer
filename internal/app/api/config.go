package api

import "github.com/LLoylento4bIx4ervei/WebServer/storage"

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LoggerLevel string `toml:"logger_level"`
	Storage     *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    "127.0.0.1:8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
