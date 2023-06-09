package configs

import (
	"gopkg.in/yaml.v2"
	"os"
)

const (
	InMemoryType   = "InMemory"
	ConcurrentType = "Concurrent"
	MySqlRepoType  = "MySql"
)

type Config struct {
	Repository    Repository     `yaml:"repository"`
	Notifications []Notification `yaml:"notifications"`
	EventBus      EventBus       `yaml:"event_bus"`
}

type Repository struct {
	Type string `yaml:"type"`
}

type Notification struct {
	Type      string   `yaml:"type"`
	Receivers []string `yaml:"receivers"`
}

type EventBus struct {
	Type             string             `yaml:"type"`
	ObserversByEvent []ObserversByEvent `yaml:"observers_by_event"`
}

type ObserversByEvent struct {
	EventName string   `yaml:"event_name"`
	Observers []string `yaml:"observers"`
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
