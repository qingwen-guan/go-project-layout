package config

import (
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

type Config struct {
	AppName string   `toml:"app_name"`
	LogDirs []string `toml:"log_dirs"`
}

func NewConfigFromFile(filepath string) (*Config, error) {
	conf := &Config{}
	if _, err := toml.DecodeFile(filepath, conf); err != nil {
		return conf, err
	}

	if conf.AppName == "" {
		parts := strings.Split(os.Args[0], "/")
		conf.AppName = parts[len(parts)-1]
	}
	return conf, nil
}
