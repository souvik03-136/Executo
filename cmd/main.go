package main

import (
	"fmt"
	"os"

	"github.com/souvik03-136/Executo/internal/logger"
	"github.com/souvik03-136/Executo/pkg/cli"
)

func main() {
	// Initialize logger
	err := logger.Init("executo.log")
	if err != nil {
		fmt.Printf("Error initializing logger: %v\n", err)
		os.Exit(1)
	}

	// Execute CLI
	if err := cli.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
