package utils

import "os"

type Config struct {
	RemoteProxy string
	ClientId    string
	AppKey      string
}

func NewConfig() *Config {
	if os.Getenv("SERVER_ADDRESS") == "" {
		os.Setenv("SERVER_ADDRESS", "127.0.0.1:8099")
	}
	if os.Getenv("CLIENT_ID") == "" {
		os.Setenv("CLIENT_ID", "debugger-client")
	}
	return &Config{
		RemoteProxy: os.Getenv("SERVER_ADDRESS"),
		ClientId:    os.Getenv("CLIENT_ID"),
		AppKey:      os.Getenv("APP_KEY"),
	}
}
