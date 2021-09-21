package utils

import "os"

type Config struct {
	RemoteProxy  string
	RemoteSchema string
	ClientId     string
	AppKey       string
}

func NewConfig() *Config {
	return &Config{
		RemoteProxy:  os.Getenv("SERVER_ADDRESS"),
		ClientId:     os.Getenv("CLIENT_ID"),
		AppKey:       os.Getenv("APP_KEY"),
		RemoteSchema: os.Getenv("REMOTE_SCHEMA"),
	}
}
