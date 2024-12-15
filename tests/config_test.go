package tests

import (
	"os"
	"testing"

	"github.com/souvik03-136/Executo/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDefaultConfig(t *testing.T) {
	// Get the default config
	cfg := config.GetDefaultConfig()

	// Assert that the scan directories are not empty
	assert.NotEmpty(t, cfg.ScanDirectories, "Default config is missing scan directories")

	// Assert that the DefaultApp is empty
	assert.Empty(t, cfg.DefaultApp, "Default config's DefaultApp should be empty")

	// Assert that shortcuts are not empty
	assert.NotEmpty(t, cfg.Shortcuts, "Default config is missing shortcuts")
}

func TestSaveAndLoadConfig(t *testing.T) {
	// Create a sample configuration
	cfg := config.Config{
		ScanDirectories: []string{"/test/bin", "/test/usr/local/bin"},
		DefaultApp:      "test_app",
		Shortcuts: map[string]string{
			"test_browser": "/test/path/firefox",
			"test_editor":  "/test/path/vim",
		},
	}

	// Temporary file for testing
	testFilePath := "test_config.json"
	defer os.Remove(testFilePath) // Clean up after test

	// Save the configuration
	err := config.SaveConfig(testFilePath, cfg)
	require.NoError(t, err, "Failed to save config")

	// Load the configuration
	loadedCfg, err := config.LoadConfig(testFilePath)
	require.NoError(t, err, "Failed to load config")

	// Assert that the loaded configuration matches the saved configuration
	assert.Equal(t, len(cfg.ScanDirectories), len(loadedCfg.ScanDirectories), "ScanDirectories do not match")
	assert.Equal(t, cfg.ScanDirectories[0], loadedCfg.ScanDirectories[0], "First ScanDirectory does not match")
	assert.Equal(t, cfg.DefaultApp, loadedCfg.DefaultApp, "DefaultApp does not match")
	assert.Equal(t, len(cfg.Shortcuts), len(loadedCfg.Shortcuts), "Shortcuts length does not match")
	assert.Equal(t, cfg.Shortcuts["test_browser"], loadedCfg.Shortcuts["test_browser"], "test_browser shortcut does not match")
}

func TestLoadConfig_FileNotExist(t *testing.T) {
	// File path that does not exist
	nonExistentFilePath := "non_existent_config.json"

	// LoadConfig should return defaultConfig
	loadedCfg, err := config.LoadConfig(nonExistentFilePath)
	require.NoError(t, err, "Failed to load config for non-existent file")

	// Validate that the loaded configuration is the default
	defaultCfg := config.GetDefaultConfig()
	assert.Equal(t, len(defaultCfg.ScanDirectories), len(loadedCfg.ScanDirectories), "ScanDirectories do not match default config")
	assert.Equal(t, defaultCfg.DefaultApp, loadedCfg.DefaultApp, "DefaultApp does not match default config")
}
