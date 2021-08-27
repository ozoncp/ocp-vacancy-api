package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

var doOnce sync.Once

func init() {
	doOnce.Do(func() {
		file, err := os.Open("config.yml")
		if err != nil {
			panic("failed to open config.yml: " + err.Error())
		}
		defer file.Close()

		decoder := yaml.NewDecoder(file)
		if err = decoder.Decode(&config); err != nil {
			panic("failed to decode config.yml: " + err.Error())
		}
	})
}

var config *Config

func GetConfig() Config {
	return *config
}

type Config struct {
	BatchSize  int              `yaml:"batch_size"`
	GRPC       ServerAddress    `yaml:"grpc"`
	REST       ServerAddress    `yaml:"rest"`
	Database   ConnectionString `yaml:"database"`
	Kafka      KafkaConfig      `yaml:"kafka"`
	Prometheus ServerAddress    `yaml:"prometheus"`
	Jaeger     ServerAddress    `yaml:"jaeger"`
}

type ServerAddress struct {
	Address string `yaml:"address"`
}

type ConnectionString struct {
	ConnectionString string `yaml:"connection_string"`
}

type KafkaConfig struct {
	Topic   string   `yaml:"topic"`
	Brokers []string `yaml:"brokers"`
}
