package db

import "github.com/kelseyhightower/envconfig"

type Config struct {
	BaseStore      string `default:"keyvalue"`
	Concurrency    string `default:"none"`
	NumberOfShards int    `default:"64"`
}

func newConfig() Config {
	var c Config
	envconfig.Process("RGDB", &c)
	return c
}

var RGDBConfig = newConfig()
