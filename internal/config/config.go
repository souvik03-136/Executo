package config

import (
	"encoding/json"
	"os"
)

// Config represents the configuration for the application.
type Config struct {
	ScanDirectories []string          `json:"scan_directories"`
	DefaultApp      string            `json:"default_app"`
	Shortcuts       map[string]string `json:"shortcuts"` // Alias to app path mapping
}

var defaultConfig = Config{
	ScanDirectories: []string{"/usr/bin", "/usr/local/bin", "/Applications"},
	DefaultApp:      "",
	Shortcuts:       make(map[string]string),
}

// LoadConfig loads the configuration from a file.
func LoadConfig(filePath string) (Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return defaultConfig, nil
		}
		return Config{}, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// SaveConfig saves the configuration to a file.
func SaveConfig(filePath string, cfg Config) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(cfg); err != nil {
		return err
	}

	return nil
}
