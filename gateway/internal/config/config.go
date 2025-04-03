package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	APIEndpoint  string        `mapstructure:"inference_api_url"`
	Port         string        `mapstructure:"port"`
	RateInterval time.Duration `mapstructure:"rate_limit_interval"`
	RateBurst    int           `mapstructure:"rate_limit_burst"`
	RateTTL      time.Duration `mapstructure:"rate_limit_ttl"`
	RateCleanup  time.Duration `mapstructure:"rate_limit_cleanup"`
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetEnvPrefix("GATEWAY")
	viper.AutomaticEnv()

	viper.SetDefault("inference_api_url", "http://127.0.0.1:8000/predict")
	viper.SetDefault("port", "8020")
	viper.SetDefault("rate_limit_interval", "6s")
	viper.SetDefault("rate_limit_burst", 5)
	viper.SetDefault("rate_limit_ttl", "10m")
	viper.SetDefault("rate_limit_cleanup", "5m")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	// Validation
	if cfg.APIEndpoint == "" {
		return nil, errors.New("api endpoint must be set")
	}
	if cfg.RateBurst <= 0 {
		return nil, errors.New("rate burst must be > 0")
	}
	if cfg.RateInterval <= 0 || cfg.RateTTL <= 0 || cfg.RateCleanup <= 0 {
		return nil, errors.New("invalid rate limiter durations")
	}

	return &cfg, nil
}
