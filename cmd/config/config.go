package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// Config manages application configuration
type Config struct {
	Home     string
	ConfigDir string
}

// New creates a new Config instance
func New() *Config {
	home, _ := os.UserHomeDir()
	configDir := filepath.Join(home, ".copilot-cli")
	return &Config{
		Home:      home,
		ConfigDir: configDir,
	}
}

// Load loads configuration from disk
func (c *Config) Load() error {
	if _, err := os.Stat(c.ConfigDir); os.IsNotExist(err) {
		return os.MkdirAll(c.ConfigDir, 0755)
	}
	return nil
}

// Save saves configuration to disk
func (c *Config) Save() error {
	fmt.Printf("Saving config to %s\n", c.ConfigDir)
	return nil
}
