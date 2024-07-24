package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"os"
)

func New(config []byte, envFile string) (*Config, error) {
	this := Config{}
	err := yaml.Unmarshal(config, &this)
	if err != nil {
		return nil, fmt.Errorf("yaml Unmarshal error: %s", err)
	}
	err = godotenv.Load(envFile)
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %s", err)
	}
	getFromEnv(&this)

	return &this, nil
}

func getFromEnv(cfg *Config) {
	if val, exists := os.LookupEnv("DB_USER"); exists {
		cfg.DbConnection.User = val
	}
	if val, exists := os.LookupEnv("DB_PASS"); exists {
		cfg.DbConnection.Pass = val
	}
	if val, exists := os.LookupEnv("SESSION_KET"); exists {
		cfg.Session.Key = val
	}
	if val, exists := os.LookupEnv("SESSION_NAME"); exists {
		cfg.Session.Name = val
	}
	if val, exists := os.LookupEnv("GATEWAY_API_KEY"); exists {
		cfg.Gateway.ApiKey = val
	}
	if val, exists := os.LookupEnv("NOTIFIER_API_KEY"); exists {
		cfg.Notifier.APIKey = val
	}
}
