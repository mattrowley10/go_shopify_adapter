package config

import (
	"fmt"
	"os"
)

type Config struct {
	Shopify Shopify
	Server  ServerConfig
}

type Shopify struct {
	BaseUrl    string
	GraphQLUrl string
	APIKey     string
	APISecret  string
}

type ServerConfig struct {
	ListenAddr string
}

func LoadEnv() (*Config, error) {
	cfg := &Config{}

	cfg.Shopify.BaseUrl = getEnvOrDefault("SHOPIFY_GRAPHQL_URL", "https://shopify.dev.com")
	cfg.Shopify.BaseUrl = getEnvOrDefault("SHOPIFY_BASE_URL", "https://shopify.dev.com")
	cfg.Shopify.APIKey = getEnvOrWarn("SHOPIFY_API_KEY")
	cfg.Shopify.APISecret = getEnvOrWarn("SHOPIFY_API_SECRET")
	cfg.Server.ListenAddr = getEnvOrDefault("LISTEN_ADDR", ":8000")

	return cfg, nil
}

func getEnvOrDefault(k, d string) string {
	if value := os.Getenv(k); value != "" {
		return value
	}
	return d
}

func getEnvOrWarn(k string) string {
	if value := os.Getenv(k); value != "" {
		return value
	} else {
		return fmt.Sprintf("%s is missing", k)
	}
}
