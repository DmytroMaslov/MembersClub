package config

import "os"

var LISTEN_PORT = "8080"

type config struct {
	Port string
}

func Load() *config {
	port, exist := os.LookupEnv("PORT")
	if !exist {
		port = LISTEN_PORT
	}
	return &config{
		Port: port,
	}
}
