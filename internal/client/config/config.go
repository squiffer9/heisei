package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Client ClientConfig `yaml:"client"`
}

type ClientConfig struct {
	ServerURL  string           `yaml:"server_url"`
	UI         UIConfig         `yaml:"ui"`
	Connection ConnectionConfig `yaml:"connection"`
}

type UIConfig struct {
	Language       string        `yaml:"language"`
	RefreshRate    time.Duration `yaml:"refresh_rate"`
	MaxThreadsShow int           `yaml:"max_threads"`
}

type ConnectionConfig struct {
	Timeout       time.Duration `yaml:"timeout"`
	RetryAttempts int           `yaml:"retry_attempts"`
}

func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}

	// Read the config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	// Parse the YAML
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	// Override with environment variables
	config.overrideWithEnv()

	// Validate the config
	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return config, nil
}

func (c *Config) overrideWithEnv() {
	if serverURL := os.Getenv("SERVER_URL"); serverURL != "" {
		c.Client.ServerURL = serverURL
	}
	if language := os.Getenv("UI_LANGUAGE"); language != "" {
		c.Client.UI.Language = language
	}
	if refreshRate := os.Getenv("UI_REFRESH_RATE"); refreshRate != "" {
		if r, err := strconv.Atoi(refreshRate); err == nil {
			c.Client.UI.RefreshRate = time.Duration(r) * time.Second
		}
	}
	if maxThreads := os.Getenv("UI_MAX_THREADS"); maxThreads != "" {
		if m, err := strconv.Atoi(maxThreads); err == nil {
			c.Client.UI.MaxThreadsShow = m
		}
	}
	if timeout := os.Getenv("CONNECTION_TIMEOUT"); timeout != "" {
		if t, err := strconv.Atoi(timeout); err == nil {
			c.Client.Connection.Timeout = time.Duration(t) * time.Second
		}
	}
	if retryAttempts := os.Getenv("CONNECTION_RETRY_ATTEMPTS"); retryAttempts != "" {
		if r, err := strconv.Atoi(retryAttempts); err == nil {
			c.Client.Connection.RetryAttempts = r
		}
	}
}

func (c *Config) validate() error {
	if c.Client.ServerURL == "" {
		return fmt.Errorf("server URL is required")
	}
	if c.Client.UI.Language == "" {
		return fmt.Errorf("UI language is required")
	}
	if c.Client.UI.RefreshRate <= 0 {
		return fmt.Errorf("invalid UI refresh rate: %v", c.Client.UI.RefreshRate)
	}
	if c.Client.UI.MaxThreadsShow <= 0 {
		return fmt.Errorf("invalid max threads to show: %d", c.Client.UI.MaxThreadsShow)
	}
	if c.Client.Connection.Timeout <= 0 {
		return fmt.Errorf("invalid connection timeout: %v", c.Client.Connection.Timeout)
	}
	if c.Client.Connection.RetryAttempts < 0 {
		return fmt.Errorf("invalid connection retry attempts: %d", c.Client.Connection.RetryAttempts)
	}
	return nil
}
