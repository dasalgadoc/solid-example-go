package configs

import (
	"gopkg.in/yaml.v2"
	"os"
)

const (
	InMemoryRepoType = "InMemory"
	MySqlRepoType    = "MySql"
)

type Config struct {
	Repository Repository `yaml:"repository"`
}

type Repository struct {
	Type string `yaml:"type"`
}

func LoadConfig(file string) (Config, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
