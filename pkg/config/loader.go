package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

const (
	env      = "ENVIRONMENT"
	location = "./config/"
)

func LoadConfig() (*Config, error) {
	if os.Getenv("RENDER") == "" {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("error loading .env file: %w", err)
		}
	}

	config := &Config{}
	configPath := location + os.Getenv(env) + ".yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file %s does not exist", configPath)
	}

	err := cleanenv.ReadConfig(configPath, config)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	return config, nil
}
