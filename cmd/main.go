package main

import (
	"fmt"
	"os"

	"github.com/souvik03-136/Executo/internal/logger"
	"github.com/souvik03-136/Executo/pkg/cli"
	"github.com/souvik03-136/Executo/ui/assets"
)

func main() {
	// Initialize logger
	err := logger.Init("executo.log")
	if err != nil {
		fmt.Printf("Error initializing logger: %v\n", err)
		os.Exit(1)
	}

	// Execute CLI commands
	if err := cli.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Initialize UI with default configurations
	config := assets.DefaultConfig() // Get default config
	config.Theme = "dark"            // Set theme to dark
	config.Layout = "grid"           // Choose grid layout
	assets.InitializeUI(config)      // Initialize UI
}
