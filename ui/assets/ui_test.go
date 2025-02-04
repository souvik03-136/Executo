package assets

import (
	"os"
	"os/exec"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()
	if config.Theme == "" || config.Layout == "" {
		t.Error("DefaultConfig returned invalid configuration")
	}
}

func TestInitializeUI(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping UI tests in short mode")
	}

	// Re-run this test as a separate process to ensure it's on the main goroutine
	if os.Getenv("RUN_MAIN") == "1" {
		config := DefaultConfig()
		InitializeUI(config) // This will now run in a new process's main goroutine
		return
	}

	// Start a separate process for UI testing
	cmd := exec.Command(os.Args[0], "-test.run=TestInitializeUI")
	cmd.Env = append(os.Environ(), "RUN_MAIN=1")
	err := cmd.Run()

	if err != nil {
		t.Fatalf("TestInitializeUI failed: %v", err)
	}
}
