package utils

import "os"

type Config struct {
	RemoteProxy string
	ClientId    string
}

func NewConfig() *Config {
	return &Config{
		RemoteProxy: os.Getenv("SERVER_ADDRESS"),
		ClientId:    os.Getenv("CLIENT_ID"),
	}
}
